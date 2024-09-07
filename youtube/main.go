package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
)

// time to live to update the health of the microservice
const ttl = time.Second * 8
const checkID = "check_health" // id used to unique identify the health check in Consul

type Service struct {
	consulClient *api.Client
}

func NewService() *Service {
	// initializes a new Consul API client
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal("Error in NewService", err)
	}
	
	// assigns the created Consul API client to the Service struct
	return &Service{
		consulClient: client,
	}
}

func (s *Service) Start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Error starting", err)
	}

	s.registerServiceToConsul() // registers the service to Consul
	go s.updateHealthCheck() // starts a goroutine to continuously update health status
	s.acceptLoop(ln) // starts the loop to accept incoming connections
}

func (s *Service) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept() // accepts incoming connections
		if err != nil {
			log.Fatal("Accept failed", err)
		}
	}
}

// inform consul before ttl that our service is running and healthy
func (s *Service) updateHealthCheck() {
	// creates a ticker for every 5 seconds
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := s.consulClient.Agent().UpdateTTL(checkID, "online", api.HealthPassing)
		if err != nil {
			log.Fatal("Error updating TTL", err)
		}

		// send a message to the channel every 5 seconds from the ticker
		<-ticker.C
	}
}

func (s *Service) registerServiceToConsul() {
	// Define a health check configuration for the services
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(), // set the duration before Consul considers the service critical if it fails to update its TTL
		TLSSkipVerify:                  true,         // skip TLS certificate verification for checks
		TTL:                            ttl.String(), // set the TTL (Time To Live) for the service
		CheckID:                        checkID,      // unique ID for the health check
	}

	// define service registration details
	register := &api.AgentServiceRegistration{
		ID:      "login_service",   // unique ID for the service registration
		Name:    "mycluster",       // name of the service
		Tags:    []string{"login"}, // tags associated with the service
		Address: "127.0.0.1", // ip address where the service is available
		Port:    3000, // port on which the service is available
		Check:   check, // set the health check for the service
	}

	// set up a watcher to monitor changes in services
	query := map[string]any{ 
		"type":        "service", // query type for the watcher
		"service":     "mycluster", // service name to watch for changes
		"passingonly": true, // watch only passing services
	}

	// create a watch plan
	// watch plan allows users to monitor changes within Consul and react to those changes
	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal("error parsing watcher", err)
	}
	
	// define a handler to process updates received from the watch plan
	plan.HybridHandler = func(index watch.BlockingParamVal, result interface{}) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			// process the received service entries
			for _, entry := range msg {
				fmt.Println("new entry joined: ", entry.Service)
			}
		}
		fmt.Println("update cluster", result)
	}

	// run the watch plan asynchronously
	go func() {
		plan.RunWithConfig("", &api.Config{})
	}()

	// register the service to Consul using the Consul client
	err = s.consulClient.Agent().ServiceRegister(register)
	if err != nil {
		log.Fatal("error registering consul client", err)
	}
}

func main() {
	s := NewService()
	s.Start()
}

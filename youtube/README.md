# Golang Consul

## Installing Consul on Mac

```
brew tap hashicorp/tap
brew install hashicorp/tap/consul

consul agent -dev

Go to http://localhost:8500

// For golang package installation
go get github.com/hashicorp/consul/api
```

## Register Service

- api.AgentServiceCheck
    - DeregisterCriticalServiceAfter
    - TLSSkipVerify
    - TTL
    - CheckID
- api.AgentServiceRegistration
    - ID
    - Name
    - Tags
    - Address
    - Port
    - Check
- c.Agent().UpdateTTL(check.CheckID, reason, api.HealthPassing)

## Watcher

- query (type, service=clustername, passinonly=true)

watch.Parse(query)

plan.HybridHandler = func(index watch.BlockingParamVal, result interface{})
plan.RunWithConfig("", &api.Config{}) => blocking
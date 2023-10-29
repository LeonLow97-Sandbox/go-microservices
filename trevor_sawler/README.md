## Microservices

- Monolithic applications: all of your business logic, connection to database, authentication, sending email or logging, is built into a single application.
- Distributed applications: instead of having 1 big monolith, you have many smaller applications that do 1 thing and do it well.
- Microservices, also known as the microservice architecture, are an architectural style which structures an application as a loosely coupled collection of smaller applications. The microservice architecture allows for the rapid and reliable delivery of large, complex applications.
  - Microservices
    - Maintainable and Testable
    - Breaking monolith up from functions/packages to completely separate programs.
    - Communicate via JSON/REST, RPC, gRPC, and over a messaging queue
    - Easy to scale and can be deployed independently.
    - Organized around business capabilities
    - Often owned by a small team

### Monolithic Architecture

<img src="./diagrams/monolithic-project.png" />

### Project Specifications

<img src="./diagrams/microservice-project.png" />

- A frontend web application that connects to 5 microservices:
  - **Broker**: optional single point of entry into the microservice cluster (microservices). (In Docker image)
  - **Authentication**: with a Postgres database
  - **Logger**: with a MongoDB database
  - **Mail**: takes a JSON payload, converts into a formatted email, and send it out
  - **Listener**: receives messages from RabbitMQ and acts upon them
- Communicate from between Microservices using:
  - REST API with JSON as transport
  - Sending & Receiving using RPC
  - Sending & Receiving using gRPC
  - Initiating and responding to events using Advanced Message Queuing Protocol (AMQP) with RabbitMQ

### Authentication Microservice

<img src="./diagrams/authentication-microservice.png" />

- Adding an Authentication Microservice that will be called by the Broker Service.
- There is a link between the browser and the Authentication Microservice (not a common practice and usually it will be called by the Broker Service).

### Logger Microservice

- Logger Service has no connection to Internet and is only available within the Microservice Cluster. This cluster can be
  - Docker
  - Docker Swarm
  - Kubernetes Cluster
- Logger Service stores all its information in Mongo (NoSQL) Database.
- Added a logging service in authentication microservice
  - Whenever someone logs in successfully or logout, it will create a log entry.

### Helpful Links

- [Relation table not created in Docker container](https://rajyavardhan.medium.com/when-you-get-relation-does-not-exist-in-postgres-7ffb0c3c674b)
- [StackOverflow - Slow Queries between server and database in Docker](https://stackoverflow.com/questions/65875996/very-slow-queries-between-server-and-database-in-docker)
- [Bind Mounts vs Volumes](https://docs.docker.com/storage/volumes/)
- [Blog on Docker volumes vs Bind mounts](https://blog.logrocket.com/docker-volumes-vs-bind-mounts/)

### Mail Service

- In this project (development), we are allowing the user to press "Test Mail" and send an email to the mail server (open to the Internet, Bad!).
- In Production, Service should not communicate with Internet.
  - For example, if you want to send an email when someone unsuccessfully logs on to the system, broker will communicate with authentication service. The authentication service will then talk to the mail service saying the user is not authenticated and then send out an email. Every microservice that needs to send an email will communicate with the mail microservice. The broker service (directly connected to user's browser/internet), is not allowed to communicate directly to the mail microservice, for security purposes and prevention of spam mails.
  - Put inside Docker Swarm or Kubernetes Cluster
- Create a Mail Server (MailHog - for development)

### Listener Service (RabbitMQ - AMQP)

<img src="./diagrams/listener-service.png" />

- Listener Service that talks to RabbitMQ (AMQP)
- If someone wants to authenticate and sends a request to the broker, the broker doesn't communicate directly with the authentication service.
  - The Broker pushes the message to RabbitMQ (AMQP Server)
  - Listener pulls a message out of the queue and calls the appropriate service based on the content in the message.
  - Listener then sends a request to the authentication service and attempts the login.
- E.g., Request --> Broker Service (publisher) --> RabbitMQ --> Listener Service (subscriber) --> Log/Authentication Microservice

### RPC

<img src="./diagrams/rpc-communication.png" />

- Communication between services must be in the same programming language, e.g., Go.
  - If broker service uses RPC in Python and Logger uses Go, it won't work.
- RPC can have better performance because it is faster than marshaling and un-marshaling JSON.
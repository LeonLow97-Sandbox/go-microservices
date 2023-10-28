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

## Helpful Links

- [Relation table not created in Docker container](https://rajyavardhan.medium.com/when-you-get-relation-does-not-exist-in-postgres-7ffb0c3c674b)
- [StackOverflow - Slow Queries between server and database in Docker](https://stackoverflow.com/questions/65875996/very-slow-queries-between-server-and-database-in-docker)
- [Bind Mounts vs Volumes](https://docs.docker.com/storage/volumes/)
- [Blog on Docker volumes vs Bind mounts](https://blog.logrocket.com/docker-volumes-vs-bind-mounts/)

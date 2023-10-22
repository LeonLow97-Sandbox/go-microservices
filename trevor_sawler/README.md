## Microservices

- Monolithic applications: all of your business logic, connection to database, authentication, sending email or logging, is built into a single application.
- Distributed applications: instead of having 1 big monolith, you have many smaller applications that do 1 thing and do it well.
    - Microservices
        - In modern applications that require scaling.
        - Breaking monolith up from functions/packages to completely separate programs.
        - Communicate via JSON/REST, RPC, gRPC, and over a messaging queue
        - Easier to scale
        - Easier to maintain
        - Harder to write

## What we will Build

- A frontend web application that connects to 5 microservices:
    - **Broker**: optional single point of entry to microservices. (In Docker image)
    - **Authentication**: Postgres
    - **Logger**: MongoDB
    - **Mail**: sends email with a specific template
    - **Listener**: consumes messages in RabbitMQ and initiates a process
- Communicate from between Microservices using:
    - REST API with JSON as transport
    - Sending & Receiving using RPC
    - Sending & Receiving using gRPC
    - Initiating and responding to events using Advanced Message Queuing Protocol (AMQP) with RabbitMQ
# Microservices

This repository serves as a practical example and learning resource for understanding microservices architecture.

The examples within this readme file are expanded upon within the project files which implements a Fibonacci number calculator using a number of different programming language.

Microservices are generally part of a whole, this repository has examples of a single Microservice and does not explore service inter-communication.

> [!Note]
> This is intended to be a grassroots introduction and as such, topics including api-gateway, service-discovery, circuit-breakers, logging, tracing, hosting/orchestration environments and other Microservice patterns are not covered.

### Services

1. C# Fibonacci Service (cs-fib)
2. Go Fibonacci Service (go-fib)
3. JavaScript Fibonacci Service (js-fib)

Each service is containerized and can be run independently. Within each, the business logic is separate from the delivery medium meaning that the core functionality can be easily adapted to different interfaces or protocols such as HTTP, gRPC, or even a CLI. If we were to expand these services and incorporate a dependency such as a database or API, hexagonal architecture (also known as ports and adapters architecture) would be an excellent next step.

You'll also notice that when building each of these services we embed a version number for the service based upon the current DVCS tag and SHA i.e. `v0.0.0-0-g2b12387` to ensure traceable, immutable versioning. Tags should follow semantic versioning (SemVer) e.g., v1.2.3. 

Go Uses the standard build system to embed the version number, whereas C# and Javascript rely on the `gen` target in the Makefile i.e. `make gen` which is a dependency for `make image`. The Makefile will help with orchestration of various tasks such as building images and should work on a Windows laptop, provided you're using an appropriate environment like Git Bash or WSL that supports the required shell commands and syntax.

## What are Microservices?

A Microservice architecture is an architectural style used to build an application that follows a modular and decoupled approach. Microservices increase flexibility and scalability by breaking down an application in to a collection of small independent services, each focusing on a specific task. Each component in a microservices architecture:
- Can be deployed independently
- Runs in its own process
- Focuses on doing one thing well

## Why do we use Microservices?

1. **Scalability**: Microservices enable you to scale individual services both vertically (allocating more resources to an instance) and horizontally (adding more instances). For instance, you can scale the payroll processing service independently of the user authentication service.
2. **Resilience**: Microservices reduce the impact of a failure. If one service goes down–when architected correctly–it shouldn’t bring down the entire application.
3. **Faster Development** and Deployment: Teams can work on and deploy different services independently, enabling parallel development, accelerating release times.
4. **Easier Maintenance** and Updates: Microservices are smaller and focused on a single task, they’re easier to understand, test, and modify.  Helping with long-term maintenance and adaptability.

Microservices also introduce complexity in managing systems, which is why tools like containerisation (Docker), orchestration (Kubernetes K8S or ECS), and DevOps practices are often paired with them.

## How do we decide what to put in a Microservice?

1. **Bounded Context**: A well-defined microservice should be small enough to focus on a single specific business capability within the application domain, ensuring a single coherent responsibility with the ability to operate independently.
2. **Independent Deployability**: A microservice should be independently deployable, without affecting or depending on other microservices. 
3. **Data Ownership**: A microservice should have its own internal database or schema
4. **Loose Coupling**: Microservices should communicate with each other through well-defined interface with no underlying knowledge of another Microservice's internal business logic or data storage. 

Good practice is to identify a specific business capability (e.g., payroll processing, user authentication, or absence management) and build a service that focuses solely on that. Dependencies that support or complement this functionality can be provided by a separate microservices.

You might want to consider how frequently each business domain changes or scales. Independent, unrelated functions with distinct lifecycles or scaling needs are often best managed as separate services.

It is often said that a team developing a microservice should be small enough to be fed by two pizzas, and that microservices should ideally have a development cycle of about two weeks.

## How do we make a Microservice

A Microservice can be written in any language that can be containerised. Here is an example of a microservice written in [Go](https://go.dev/);

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    http.HandleFunc("/products", productsHandler)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    products := []Product{
        {ID: 1, Name: "Product A", Price: 9.99},
        {ID: 2, Name: "Product B", Price: 14.99},
        {ID: 3, Name: "Product C", Price: 19.99},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```

To run this microservice, save the code to a file (e.g., `main.go`) and execute the following command:

```sh
go run main.go
```

Then, you can send a GET request to `http://localhost:8080/products` to see the list of products.

> [!Note]
> In a real-world scenario, you'd likely want to add more functionality, such as authentication, authorisation, connection to a datastore, logging, error handling, and testing.

## How do we prepare a Microservice for deployment

To deploy a Microservice, we need to containerise it (with Docker). This enables us to package our application, its dependencies, and configuration into a standardised unit that can run consistently across different environments. Containerisation provides a layer of abstraction between the application and the underlying operating system, ensuring that the service behaves the same regardless of where it's deployed.

1. **Create a Docker file**: Here’s an example Dockerfile for the Go microservice described above

```Dockerfile
# Use an official Go image as a build environment
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o products-service .

# Use a smaller image for the final container
FROM alpine:3.18

# Copy the binary from the builder
COPY --from=builder /app/products-service .

# Use an unprivileged user.
USER nobody:nobody

# Set the container entry point
ENTRYPOINT ["/products-service"]

```

2. **Build the Docker Image**: using the following command

```sh
docker build -t products-service .
```

You can run the Docker container locally to test that everything worked as expected

```sh
docker run -p 8080:8080 products-service
```

> [!Note]
> In a real-world scenario, you'd likely want to optimise and harden the container for production, using a distro-less base image, non-root user and other best practices for security

## How do we deploy a Microservice?
When deploying a service, a container orchestration system such as Kubernetes (K8S) or AWS Elastic Container Service (ECS) will pull the service image from a container registry (e.g., Docker Hub, AWS ECR, or Google Container Registry) so that it can be deployed to a production environment. This means that in order to deploy a microservice the service image should first be uploaded/pushed into a container registry.

In order to push a service image to a container registry it will need a tag containing the registry URI, image name, and versioning tag

```sh
docker tag products-service $(registry)/products-service:$(revision)
```

The docker command can then be used to push the image to the registry

```sh
docker push $(registry)/products-service:$(revision)
```

A service should be deployed using continuous integration and continuous deployment (CI/CD) pipeline to automatically build, test and deploy new versions of the microservice in response to something like a commit to a distributed version control system such as git or the acceptance of a pull request.

Configuration for the microservice, such as, environment variables, rules for autoscaling, load balancing and monitoring are added to these CI/CD pipeline scripts to ensure that the microservice executes in the correct context and remain healthy.

## What other things should we consider?

- **Stateless**: In order to support load balancing and autoscaling microservices should be designed to be stateless meaning that they shouldn't store any client specific session information.
- **Correlation IDs**: Adding correlation identifiers to service requests enable us to trace requests cost multiple services.
- **Authorisation**: JSON Web Tokens (JWT) are an ideal mechanism for enabling find grained permissions.
- **Authentication**: 
- **Secrets Management**: 
- **Asynchronicity**: Consider using asynchronous non-blocking communication patterns such as message queues to decouple services and improve resilience.
- **Protocols**: While HTTP is a straightforward choice for microservices, gRPC can offer advantages, especially for high-performance, inter-service communication.
- **Observability**:

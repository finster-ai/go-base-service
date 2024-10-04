# go-base-service


## Description

This is a the base template repo of Go microservices that provides this base examples:
- REST API
- GRPC
- Kafka consumer and Producer
- MongoDB connection
- Redis connection

The idea behind this repo is to provide a base for a microservice that can be easily clone and extended to provide more functionality.

## Folder Structure

The folder structure is as follows, trying to stick to the most standard Go project structure:

```go
/go-base-service
│
├── /api               # API-specific code (handlers, routes)
│   ├── /graphql
│   │   ├── schema.go       # GraphQL schema definitions
│   │   ├── resolvers.go    # Resolvers to handle query/mutation logic
│   │   └── graphql_handler.go  # Entry point for GraphQL handling (e.g., setup of the GraphQL server)
│   ├── /grpc
│   │   ├── /pb            # Protobuf-generated code
│   │   │   ├── user.pb.go  # Generated code from user.proto
│   │   │   ├── product.pb.go # Generated code from product.proto
│   │   ├── user_handler.go # gRPC service implementation for user
│   │   ├── product_handler.go  # gRPC service implementation for product
│   │   └── grpc_server.go  # gRPC server setup
│   └── /http
│       ├── /request        # Request structures
│       │   ├── user_request.go
│       │   └── product_request.go
│       ├── /response       # Response structures
│       │   ├── user_response.go
│       │   └── product_response.go
│       ├── user_handler.go # Handler for user-related HTTP endpoints
│       ├── product_handler.go  # Handler for product-related HTTP endpoints
│       └── router.go       # HTTP router setup, defining routes for all endpoints
│
├── /charts            # Helm charts for Kubernetes deployment
│
├── /cmd
│   └── /main.go       # Entry point for microservice (main.go)
│
├── /internal
│   ├── /cache         # Redis client and cache-related logic
│   ├── /kafka         # Kafka consumer and producer logic
│   ├── /model         # Business Data models (MongoDB collections, Redis structures)
│   ├── /repository    # Data access logic (MongoDB, Redis, other databases)
│   ├── /server        # HTTP and gRPC server setup
│   └── /service       # Business logic and use case implementations
│
├── /pkg               # Shared code across services or other Go projects
│   ├── /config        # Configuration-related code (environment variables, configs)
│   ├── /logging       # Logging setup
│   ├── /middleware    # Middlewares for HTTP/gRPC requests
│   └── /utils         # Utility functions/helpers
│
├── /proto             # .proto files (gRPC service definitions)
│
├── /scripts           # Deployment and maintenance scripts (e.g., for database migrations)
│
├── /test              # Test files (unit, integration, etc.)
│
├── Dockerfile         # Dockerfile for containerization
├── go.mod             # Go module dependencies
└── go.sum             # Go module dependency lock file
```

Once you've have created your repo using this one as a template, you can delete unnecessary folders




## Tools, Frameworks Etc

// TODO: Add the tools, frameworks, libraries, etc that are used in this project



## Installation

// TODO: Add the installation instructions for this project IF NECESSARY


### Running locally

The backend needs access to GCP to run. Before running it for the first time, you need to authenticate running this command in your terminal.
```
gcloud auth application-default login
```
If this doesn't work, please check the documentation <https://cloud.google.com/docs/authentication/provide-credentials-adc>



### Code Structure & Basic Flow

* Endpoints are defined in [`controller.py`]
* GRPC endpoints logic is implemented in [`base_model1_grpc_impl.py`]
* The entry point of the app is in [`api_base_service.py`]
* By default the app starts with a subscriber to a pubsub topic, an REST API server and a GRPC server



### Modifying Proto Files

You need to run this command to generate new  grpc files after proto has been updated
```
python -m grpc_tools.protoc -I./app/proto --python_out=./app/proto/gen --grpc_python_out=./app/proto/gen ./app/proto/BaseModel1.proto
```




## Building Image and Running Locally
### Build Image
```
docker build --platform linux/amd64 -t python-base-service:local .
```
### Deploy Image
```
docker run --platform linux/amd64 -d -p 8080:8080 \
--name python-base-service-local \
-v /Users/emmanuelcastillodecarvalho/Finster\ AI/.gke-automation-emmanuel-daring-keep-408013.json:/var/secrets/google/key.json \
-e GOOGLE_APPLICATION_CREDENTIALS="/var/secrets/google/key.json" \
python-base-service:local
```

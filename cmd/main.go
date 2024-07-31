package main

import (
	"fmt"
	"go-base-service/grpcHandler"
	"log"

	"github.com/go-micro/plugins/v4/server/grpc"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"github.com/gorilla/mux"
	"go-base-service/controller"
	"go-base-service/graphql"
	pb "go-base-service/proto/gen"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	SERVER_NAME = "goBaseServer" // server name
)

func main() {
	// Set up the gRPC service
	grpcService := setupGRPCService()
	// Set up the HTTP API service
	apiService := setupAPIService()

	// Start GraphQL server
	go func() {
		fmt.Println("Starting GraphQL server...")
		graphql.StartGraphQLServer()
	}()

	// Start the gRPC service
	go func() {
		if err := grpcService.Run(); err != nil {
			fmt.Printf("Failed to run gRPC service: %v\n", err)
		}
	}()

	// Start the HTTP API service
	go func() {
		if err := apiService.Run(); err != nil {
			fmt.Printf("Failed to run HTTP service: %v\n", err)
		}
	}()

	//TODO como consumo de 2 colas diferentes, agarro y meto aca 2 consumers, o lo hago como en java
	//TODO que tenia un consumer y un switch en base de que queue venia el mensaje
	//// Start Kafka consumer
	//go func() {
	//	log.Println("Starting Kafka consumer...")
	//kafka.ConsumeMessages("your-topic")
	//}()

	select {}

}

//TODO metodo que tome el path y el method de una annotation en el archivo de controller y lo registre en el router
// haciendo alcomo como el comando de abajo, tipo un
// for each annotation in controller
// hacer una linea como la de abajo

// router.HandleFunc("/items/{id}", controller.GetItem).Methods("GET")

func setupGRPCService() micro.Service {

	grpcServer := grpc.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":6585"),
	)
	service := micro.NewService(
		micro.Server(grpcServer),
		micro.Name("goBaseService"),
	)
	service.Init()

	// Register GRPC Handlers
	pb.RegisterBaseModel1GRPCServiceHandler(service.Server(), grpcHandler.New())

	return service
}

func setupAPIService() micro.Service {
	// Create a new mux router for the HTTP server
	router := setupRouter()

	// Set up the HTTP service
	apiServer := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8081"),
	)
	apiService := micro.NewService(
		micro.Server(apiServer),
		micro.Registry(registry.NewRegistry()),
	)
	apiService.Init()

	// Register the custom handler with the HTTP server
	if err := apiServer.Handle(newMuxHandler("muxHandler", router)); err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}

	return apiService
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Define REST API endpoints using the controller package
	router.HandleFunc("/items/{id}", controller.GetItem).Methods("GET")
	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")

	return router
}

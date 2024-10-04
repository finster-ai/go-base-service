package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/server/grpc"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-base-service/api/graphql"
	grpc2 "go-base-service/api/grpc"
	"go-base-service/api/http"
	server2 "go-base-service/internal/server"
	"go-base-service/pkg/config"

	pb "go-base-service/proto/gen"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"log"
)

var SERVER_NAME string
var API_PORT string
var GRPC_PORT string

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	printTestLogs()

	// Load the configuration file
	config.LoadConfig("config/config_local.yaml") // Adjust the path to your config file
	SERVER_NAME = config.AppConfig.GetString("app.name")
	API_PORT = config.AppConfig.GetString("server.port")
	GRPC_PORT = config.AppConfig.GetString("grpc.port")

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

func setupGRPCService() micro.Service {

	grpcServer := grpc.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":"+GRPC_PORT),
	)
	service := micro.NewService(
		micro.Server(grpcServer),
		micro.Name("goBaseService"),
	)
	service.Init()

	// Register GRPC Handlers
	pb.RegisterBaseModel1GRPCServiceHandler(service.Server(), grpc2.New())

	return service
}

func setupAPIService() micro.Service {
	// Create a new mux router for the HTTP server
	router := setupRouter()

	// Set up the HTTP service
	apiServer := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":"+API_PORT),
	)
	apiService := micro.NewService(
		micro.Server(apiServer),
		micro.Registry(registry.NewRegistry()),
	)
	apiService.Init()

	// Register the custom handler with the HTTP server
	if err := apiServer.Handle(server2.newMuxHandler("muxHandler", router)); err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}

	return apiService
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Initialize the controller and its dependencies
	http.InitController()

	// Define REST API endpoints using the controller package
	router.HandleFunc("/items/{id}", http.GetItem).Methods("GET")
	router.HandleFunc("/items", http.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", http.DeleteItem).Methods("DELETE")

	return router
}

func printTestLogs() {
	logrus.Debug("This is a DEBUG message.")
	logrus.Info("This is an INFO message.")
	logrus.Warn("This is a WARNING message.")
	logrus.Error("This is an ERROR message.")
	logrus.WithField("level", "CRITICAL").Error("This is a CRITICAL message.")
	//logrus.Fatal("This is a CRITICAL message.") // logrus.Fatal exits the program after logging
}

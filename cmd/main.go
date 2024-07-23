package main

import (
	//aca se importan las carpetas/paquetes que se crearon, no los archivos
	"github.com/go-micro/plugins/v4/server/grpc"
	"go-base-service/graphql"
	"go-base-service/grpcHandler"
	pb "go-base-service/proto/gen"
	"go-micro.dev/v4"
	"log"
)

func main() {

	// Assuming grpcHandler.NewBaseModel1GRPCService() returns an instance of your service handler
	service := micro.NewService(
		// This needs to be first as it replaces the underlying server
		// which causes any configuration set before it
		// to be discarded
		micro.Server(grpc.NewServer()),
		micro.Name("goBaseService"),
		// Other options...
	)
	service.Init()

	// Register GRPC Handlers
	pb.RegisterBaseModel1GRPCServiceHandler(service.Server(), grpcHandler.New())

	//TODO como consumo de 2 colas diferentes, agarro y meto aca 2 consumers, o lo hago como en java
	//TODO que tenia un consumer y un switch en base de que queue venia el mensaje
	//// Start Kafka consumer
	//go func() {
	//	log.Println("Starting Kafka consumer...")
	//kafka.ConsumeMessages("your-topic")
	//}()

	// Start GraphQL server
	go func() {
		log.Println("Starting GraphQL server...")
		graphql.StartGraphQLServer()
	}()

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	select {}
}

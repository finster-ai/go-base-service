package grpc_handler

import (
	"context"
	pb "go-base-service/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BaseModel1GRPCService struct{}

func New() *BaseModel1GRPCService {
	return &BaseModel1GRPCService{}
}

func (s *BaseModel1GRPCService) ExampleCall(ctx context.Context, req *pb.ExampleCallRequest, rsp *pb.ExampleCallResponse) error {
	// Example validation (adjust according to your actual validation logic)
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	// Assuming some business logic here that might fail
	//if err := someBusinessLogic(); err != nil {
	//    // Log the error, adjust logging according to your setup
	//    log.Printf("Error in business logic: %v", err)
	//    return status.Errorf(codes.Internal, "internal error occurred")
	//}

	// Directly modify the rsp parameter instead of returning a new instance
	rsp.TransactionsFound = true
	rsp.TransactionsCount = 42

	return nil
}

func (s *BaseModel1GRPCService) ExampleCallReturnsEmpty(ctx context.Context, req *pb.ExampleCallReturnsEmptyRequest, rsp *emptypb.Empty) error {
	*rsp = emptypb.Empty{} // Directly modify the rsp parameter if needed
	return nil
}

//func (s *BaseModel1GRPCService) ExampleCall(ctx context.Context, req *pb.ExampleCallRequest) (*pb.ExampleCallResponse, error) {
//	// Example validation (adjust according to your actual validation logic)
//	if req == nil {
//		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
//	}
//
//	// Assuming some business logic here that might fail
//	//if err := someBusinessLogic(); err != nil {
//	//	// Log the error, adjust logging according to your setup
//	//	log.Printf("Error in business logic: %v", err)
//	//	return nil, status.Errorf(codes.Internal, "internal error occurred")
//	//}
//
//	return &pb.ExampleCallResponse{
//		TransactionsFound: true,
//		TransactionsCount: 42,
//	}, nil
//}

//func (s *BaseModel1GRPCService) ExampleCallReturnsEmpty(ctx context.Context, req *pb.ExampleCallReturnsEmptyRequest) (*emptypb.Empty, error) {
//	return &emptypb.Empty{}, nil
//}

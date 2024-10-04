package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "go-base-service/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
)

type BaseModel1GRPCService struct{}

func New() *BaseModel1GRPCService {
	return &BaseModel1GRPCService{}
}

func (s *BaseModel1GRPCService) ExampleCall(ctx context.Context, req *pb.ExampleCallRequest, rsp *pb.ExampleCallResponse) error {
	// Example validation (adjust according to your actual validation logic)
	logrus.Info("ExampleCall GRPC Endpoint Reached")
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	rsp.TransactionsFound = true
	rsp.TransactionsCount = int64(rand.Intn(100)) // Returns a random number between 0 and 99
	logrus.Info("ExampleCall GRPC Endpoint Process Finished - Returning Response: " + rsp.String())
	return nil
}

func (s *BaseModel1GRPCService) ExampleCallReturnsEmpty(ctx context.Context, req *pb.ExampleCallReturnsEmptyRequest, rsp *emptypb.Empty) error {
	logrus.Info("ExampleCallReturnsEmpty GRPC Endpoint Reached")
	*rsp = emptypb.Empty{} // Directly modify the rsp parameter if needed
	logrus.Info("ExampleCallReturnsEmpty GRPC Endpoint Process Finished - Returning EMPTY Response")
	return nil
}

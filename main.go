package main

import (
	"context"
	"log"
	"net" // 替换为生成的 Go 文件的包路径
	pb "test/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println(123)
	return &pb.EchoResponse{Message: "Echo: " + req.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "ArSearch/proto"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedArSearchServer
}

func (s *server) SearchOnCluster(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	//todo impl
	fmt.Println("ack")
	list :=make([]*pb.SearchResponseItem,0)
	return &pb.SearchResponse{Count: int32(len(list)),SearchItems: list,}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArSearchServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



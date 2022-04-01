package grpc

import (
	"context"
	"log"
	"time"

	pb "ArSearch/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ArSearch interface {
	SearchOnCluster(ctx context.Context, req pb.SearchRequest) (*pb.SearchResponse, error)
}

type ArSearchImpl struct{}

var DefaultArSearchImpl = &ArSearchImpl{}

func (a *ArSearchImpl) SearchOnCluster(ctx context.Context, req pb.SearchRequest) (*pb.SearchResponse, error) {
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewArSearchClient(conn)

	// Contact the server and print out its response.
	ctx1, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.SearchOnCluster(ctx1, &pb.SearchRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r, nil
}

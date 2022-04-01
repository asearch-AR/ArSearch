package grpc

import (
	"context"
	"testing"

	pb "ArSearch/proto"
)

func TestArSearchImpl_SearchOnCluster(t *testing.T) {
	ctx:=context.TODO()
	DefaultArSearchImpl.SearchOnCluster(ctx,pb.SearchRequest{})
}

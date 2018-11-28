package ingester

import (
	"context"
	"sync"

	"h12.io/msa"

	"google.golang.org/grpc"
	"h12.io/msa/proto"
)

type fakeStorageClient struct {
	err  error
	reqs []*proto.UpsertRequest
	mu   sync.RWMutex
}

func (c *fakeStorageClient) Upsert(ctx context.Context, in *proto.UpsertRequest, opts ...grpc.CallOption) (*proto.UpsertReply, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err != nil {
		return nil, c.err
	}
	c.reqs = append(c.reqs, in)
	return &proto.UpsertReply{Code: msa.ReplyOK}, nil
}

package storage

import (
	"context"

	"h12.io/exetl"
	"h12.io/exetl/proto"
)

// Service implements the storage gRPC service
type Service struct {
	db exetl.DB
}

// NewService creates a new storage.Service
func NewService(db exetl.DB) *Service {
	return &Service{db: db}
}

// Upsert handles an upsert request for upserting records into stroage service
func (s *Service) Upsert(ctx context.Context, req *proto.UpsertRequest) (*proto.UpsertReply, error) {
	records := make([]exetl.Record, 0, len(req.Records))
	for _, rec := range req.Records {
		records = append(records, rec.ToDomain())
	}
	if err := s.db.Upsert(req.Table, records); err != nil {
		return &proto.UpsertReply{
			Code: exetl.ReplyErr,
			Msg:  err.Error(),
		}, nil
	}
	return &proto.UpsertReply{Code: exetl.ReplyOK}, nil
}

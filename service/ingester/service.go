package ingester

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"h12.io/msa"

	"h12.io/msa/proto"
)

// Service implements an HTTP service that ingest CSV files and upsert into backend storage
type Service struct {
	storage   proto.StorageClient
	batchSize int
}

// NewService creates a new storage.Service
func NewService(storage proto.StorageClient, batchSize int) *Service {
	return &Service{storage: storage, batchSize: batchSize}
}

// ServeHTTP implements http.Handler
func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	table := q.Get("table")
	keyNames := make(map[string]bool)
	for _, key := range strings.Split(q.Get("keys"), ",") {
		keyNames[key] = true
	}
	/*
		TODO:
		* a parameter to select some of the fields
		* parameters to specify a clean method for a field, e.g. phone number
	*/

	if err := readCSV(req.Body, keyNames, s.batchSize, func(records []*proto.Record) error {
		return s.upsert(req.Context(), table, records)
	}); err != nil {
		if _, is := err.(*processError); is {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (s *Service) upsert(ctx context.Context, table string, records []*proto.Record) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	resp, err := s.storage.Upsert(ctx, &proto.UpsertRequest{
		Table:   table,
		Records: records,
	})
	if err != nil {
		return err
	}
	if resp.Code != msa.ReplyOK {
		return errors.New(resp.Msg)
	}
	return nil
}

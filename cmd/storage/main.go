package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	"h12.io/msa"
	"h12.io/msa/db/memdb"
	"h12.io/msa/proto"
	"h12.io/msa/service"
	"h12.io/msa/service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Host string
}

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.Host, "host", ":"+strconv.Itoa(msa.StorageDefaultPort), "host of the storage service")
	flag.Parse()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config) error {
	lis, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		return err
	}

	// TODO: use a real DB backend
	db := memdb.New()

	server := grpc.NewServer()
	proto.RegisterStorageServer(server, storage.NewService(db))
	reflection.Register(server)

	service.NotifyStop(func(context.Context) error {
		server.GracefulStop()
		return nil
	})

	return server.Serve(lis)
}

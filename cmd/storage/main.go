package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/h12w/msa"
	"github.com/h12w/msa/db/memdb"
	"github.com/h12w/msa/proto"
	"github.com/h12w/msa/service/storage"
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

	return server.Serve(lis)
}

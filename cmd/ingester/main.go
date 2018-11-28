package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/h12w/msa"
	"github.com/h12w/msa/proto"
	"github.com/h12w/msa/service/ingester"
	"google.golang.org/grpc"
)

type config struct {
	Batch   int
	Host    string
	Storage string
}

func main() {
	cfg := &config{}
	defaultStorageHost := os.Getenv("STORAGE_HOST")
	if defaultStorageHost == "" {
		defaultStorageHost = ":" + strconv.Itoa(msa.StorageDefaultPort)
	}
	flag.StringVar(&cfg.Host, "host", ":"+strconv.Itoa(msa.IngesterDefaultPort), "host of the ingester service")
	flag.StringVar(&cfg.Storage, "storage", defaultStorageHost, "host of the storage service")
	flag.IntVar(&cfg.Batch, "batch", 100, "processing batch size")
	flag.Parse()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config) error {
	storageConn, err := grpc.Dial(cfg.Storage, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer storageConn.Close()

	s := &http.Server{
		Addr:           cfg.Host,
		Handler:        ingester.NewService(proto.NewStorageClient(storageConn), cfg.Batch),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}

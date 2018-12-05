package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"h12.io/exetl"
	"h12.io/exetl/proto"
)

type config struct {
	Host   string
	Table  string
	Record proto.Record
}

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.Host, "host", "127.0.0.1:"+strconv.Itoa(exetl.StorageDefaultPort), "host of the storage service")
	flag.StringVar(&cfg.Table, "table", "test", "table to be upserted")
	flag.Parse()

	for i, arg := range flag.Args() {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			log.Fatalf("invalid argument %v, expecting key=value", arg)
		}
		field := &proto.Field{
			Name:  kv[0],
			Value: kv[1],
		}
		cfg.Record.Fields = append(cfg.Record.Fields, field)
		if i == 0 {
			cfg.Record.Key = field
		}
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config) error {
	conn, err := grpc.Dial(cfg.Host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := proto.NewStorageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Upsert(ctx, &proto.UpsertRequest{Table: cfg.Table, Records: []*proto.Record{&cfg.Record}})
	if err != nil {
		return err
	}
	if r.Code != exetl.ReplyOK {
		return errors.New(r.GetMsg())
	}
	return nil
}

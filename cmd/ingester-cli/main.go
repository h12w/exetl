package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"h12.io/exetl"
)

type config struct {
	Host  string
	Table string
	File  string
	Keys  string
}

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.Host, "host", "127.0.0.1:"+strconv.Itoa(exetl.IngesterDefaultPort), "host of the storage service")
	flag.StringVar(&cfg.Table, "table", "test", "table to be upserted")
	flag.StringVar(&cfg.File, "file", "", "filename of CSV file")
	flag.StringVar(&cfg.Keys, "keys", "", "comma separated key list of the CSV file")
	flag.Parse()

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config) error {
	cli := &http.Client{}
	f, err := os.Open(cfg.File)
	if err != nil {
		return err
	}
	defer f.Close()
	uri := &url.URL{
		Scheme: "http",
		Host:   cfg.Host,
	}
	query := make(url.Values)
	query.Set("table", cfg.Table)
	query.Set("keys", cfg.Keys)
	uri.RawQuery = query.Encode()
	resp, err := cli.Post(uri.String(), "text/csv", f)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}

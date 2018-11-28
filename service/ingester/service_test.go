package ingester

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIngester(t *testing.T) {
	storageClient := &fakeStorageClient{}
	server := httptest.NewServer(NewService(storageClient, 10))
	defer server.Close()

	cli := &http.Client{}
	uri, _ := url.Parse(server.URL)
	query := make(url.Values)
	query.Set("table", "test")
	query.Set("keys", "id")
	uri.RawQuery = query.Encode()
	resp, err := cli.Post(uri.String(), "text/csv", strings.NewReader(`
id,name
1,a
2,b
`))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expect 200 got %d", resp.StatusCode)
	}

	if upserts := len(storageClient.reqs); upserts != 1 {
		t.Fatalf("expect 1 upsert but got %d", upserts)
	}

	// TODO: check upserts values
}

package ingester

import (
	"reflect"
	"strings"
	"testing"

	"github.com/h12w/msa/proto"
)

func TestReadCSV(t *testing.T) {
	csvText := `
id name
1 a 
2 b
`
	records := []*proto.Record{
		{
			Key: &proto.Field{
				Name:  "id",
				Value: "1",
			},
			Fields: []*proto.Field{
				{
					Name:  "id",
					Value: "1",
				},
				{
					Name:  "name",
					Value: "a",
				},
			},
		},
		{
			Key: &proto.Field{
				Name:  "id",
				Value: "2",
			},
			Fields: []*proto.Field{
				{
					Name:  "id",
					Value: "2",
				},
				{
					Name:  "name",
					Value: "b",
				},
			},
		},
	}

	recs := []*proto.Record{}
	readCSV(strings.NewReader(csvText), map[string]bool{"id": true}, 10, func(records []*proto.Record) error {
		recs = append(recs, records...)
		return nil
	})
	if reflect.DeepEqual(recs, records) {
		t.Fatalf("expect %v got %v", records, recs)
	}
}

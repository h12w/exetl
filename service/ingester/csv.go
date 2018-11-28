package ingester

import (
	"encoding/csv"
	"io"
	"strings"

	"h12.io/msa/proto"
)

type processError struct {
	err error
}

func (e *processError) Error() string {
	return e.err.Error()
}

func readCSV(csvFile io.Reader, keyNames map[string]bool, batchSize int, process func([]*proto.Record) error) error {
	rd := csv.NewReader(csvFile)

	header, err := rd.Read()
	if err != nil {
		return err
	}
	isKey := make(map[int]bool)
	for i, fieldName := range header {
		if keyNames[fieldName] {
			isKey[i] = true
		}
	}

	records := make([]*proto.Record, 0, batchSize)
	for {
		csvRecord, err := rd.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if len(csvRecord) != len(header) {
			continue // skip misaligned record
		}

		keys := []string{}
		fields := make([]*proto.Field, 0, len(csvRecord))
		for i, value := range csvRecord {
			if isKey[i] {
				keys = append(keys, value)
			}
			fields = append(fields, &proto.Field{
				Name:  header[i],
				Value: value,
			})
		}
		key := strings.Join(keys, "+")

		records = append(records, &proto.Record{
			Key: &proto.Field{
				Name:  "key",
				Value: key,
			},
			Fields: fields,
		})

		// send by batch
		if len(records) >= batchSize {
			if err := process(records); err != nil {
				return &processError{err: err}
			}
			records = records[:0]
		}
	}

	// send the rest
	if len(records) > 0 {
		if err := process(records); err != nil {
			return &processError{err: err}
		}
	}

	return nil
}

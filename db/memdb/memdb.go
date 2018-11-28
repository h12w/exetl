// Package memdb provides a pure memory implementation of msa.DB interface
package memdb

import (
	"fmt"
	"log"
	"sync"

	"h12.io/msa"
)

type (
	// MemDB is a pure memory implementation of msa.DB interface
	MemDB struct {
		tables map[string]*table
		mu     sync.RWMutex
	}

	table struct {
		name    string
		records map[string]msa.Record
		mu      sync.RWMutex
	}
)

// New creates a new mem db
func New() *MemDB {
	return &MemDB{
		tables: make(map[string]*table),
	}
}

// Upsert satisfies msa.DB.Upsert
func (db *MemDB) Upsert(tableName string, records []msa.Record) error {
	db.mu.Lock()
	table, ok := db.tables[tableName]
	if !ok {
		table = newTable(tableName)
		db.tables[tableName] = table
	}
	db.mu.Unlock()
	return table.upsert(records)
}

func newTable(name string) *table {
	return &table{
		name:    name,
		records: make(map[string]msa.Record),
	}
}

func (t *table) upsert(records []msa.Record) error {
	t.mu.Lock()
	for _, record := range records {
		t.records[fmt.Sprint(record.Key.Value)] = record

		// TODO: just for illustration, to be replaced with a debug level logger
		log.Printf("inserted %v %+v", record.Key.Value, record)
	}
	t.mu.Unlock()
	return nil
}

package exetl

// DB related type definitions
type (
	// DB is an interface that abstracts common operations of backend storage
	DB interface {
		Upsert(table string, records []Record) error
	}

	// Record represents a DB record
	Record struct {
		Key    Field
		Fields []Field
	}

	// Field represents a field in a record
	Field struct {
		Name  string
		Value interface{}
	}
)

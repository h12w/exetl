package proto

import "h12.io/exetl"

// ToDomain converts from proto buffer type to domain type
func (r *Record) ToDomain() exetl.Record {
	key := r.GetKey()
	fields := make([]exetl.Field, 0, len(r.Fields))
	for _, field := range r.Fields {
		fields = append(fields, field.ToDomain())
	}
	return exetl.Record{
		Key: exetl.Field{
			Name:  key.Name,
			Value: key.Value,
		},
		Fields: fields,
	}
}

// ToDomain converts from proto buffer type to domain type
func (f *Field) ToDomain() exetl.Field {
	return exetl.Field{
		Name:  f.GetName(),
		Value: f.GetValue(),
	}
}

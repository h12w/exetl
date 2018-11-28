package proto

import "h12.io/msa"

// ToDomain converts from proto buffer type to domain type
func (r *Record) ToDomain() msa.Record {
	key := r.GetKey()
	fields := make([]msa.Field, 0, len(r.Fields))
	for _, field := range r.Fields {
		fields = append(fields, field.ToDomain())
	}
	return msa.Record{
		Key: msa.Field{
			Name:  key.Name,
			Value: key.Value,
		},
		Fields: fields,
	}
}

// ToDomain converts from proto buffer type to domain type
func (f *Field) ToDomain() msa.Field {
	return msa.Field{
		Name:  f.GetName(),
		Value: f.GetValue(),
	}
}

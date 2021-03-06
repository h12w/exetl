package exetl

// gRPC reply code
const (
	// ReplyOK is returned when gRPC call succeeded
	ReplyOK = iota
	// ReplyErr is returned when gRPC call failed
	ReplyErr
)

// default ports of gRPC services
const (
	StorageDefaultPort  = 9100
	IngesterDefaultPort = 9101
)

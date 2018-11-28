package msa

// gRPC reply code
const (
	// ReplyOK is returned when gRPC call succeeded
	ReplyOK = iota
	// ReplyErr is returned when gRPC call failed
	ReplyErr
)

// default ports of gRPC services
const (
	StorageDefaultPort = 9000 + iota
	IngesterDefaultPort
)

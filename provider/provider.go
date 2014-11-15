package provider

import "io"

// Provider provides data to and from the source
type Provider interface {
	io.ReadWriteCloser
	io.ReaderAt
	io.WriterAt
	Location() string
	Create(id, codec, size, prefix string)
}

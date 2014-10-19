package provider

import "io"

// Provider provides data to and from the source
type Provider interface {
	io.ReadWriteCloser
	Location() string
	Create(id, codec, size, prefix string)
}

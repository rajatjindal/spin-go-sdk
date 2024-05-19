// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package network

import (
	"unsafe"
)

// IPv6AddressShape is used for storage in variant or result types.
type IPv6AddressShape struct {
	shape [unsafe.Sizeof(IPv6Address{})]byte
}

// IPv6SocketAddressShape is used for storage in variant or result types.
type IPv6SocketAddressShape struct {
	shape [unsafe.Sizeof(IPv6SocketAddress{})]byte
}

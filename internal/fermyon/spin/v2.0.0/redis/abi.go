// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package redis

import (
	"github.com/ydnar/wasm-tools-go/cm"
	"unsafe"
)

// ErrorShape is used for storage in variant or result types.
type ErrorShape struct {
	shape [unsafe.Sizeof(Error{})]byte
}

// OptionPayloadShape is used for storage in variant or result types.
type OptionPayloadShape struct {
	shape [unsafe.Sizeof(cm.Option[Payload]{})]byte
}

// PayloadShape is used for storage in variant or result types.
type PayloadShape struct {
	shape [unsafe.Sizeof(Payload{})]byte
}

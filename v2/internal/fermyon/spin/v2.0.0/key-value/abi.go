// Code generated by wit-bindgen-go. DO NOT EDIT.

package keyvalue

import (
	"github.com/ydnar/wasm-tools-go/cm"
	"unsafe"
)

// ErrorShape is used for storage in variant or result types.
type ErrorShape struct {
	shape [unsafe.Sizeof(Error{})]byte
}

// OptionListU8Shape is used for storage in variant or result types.
type OptionListU8Shape struct {
	shape [unsafe.Sizeof(cm.Option[cm.List[uint8]]{})]byte
}

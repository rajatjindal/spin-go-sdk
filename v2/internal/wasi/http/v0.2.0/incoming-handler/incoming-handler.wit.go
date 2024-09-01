// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package incominghandler represents the exported interface "wasi:http/incoming-handler@0.2.0".
//
// This interface defines a handler of incoming HTTP Requests. It should
// be exported by components which can respond to HTTP Requests.
package incominghandler

import (
	"github.com/fermyon/spin-go-sdk/v2/internal/wasi/http/v0.2.0/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

//go:wasmexport wasi:http/incoming-handler@0.2.0#handle
//export wasi:http/incoming-handler@0.2.0#handle
func wasmexport_Handle(request0 uint32, responseOut0 uint32) {
	request := cm.Reinterpret[types.IncomingRequest]((uint32)(request0))
	responseOut := cm.Reinterpret[types.ResponseOutparam]((uint32)(responseOut0))
	Exports.Handle(request, responseOut)
	return
}

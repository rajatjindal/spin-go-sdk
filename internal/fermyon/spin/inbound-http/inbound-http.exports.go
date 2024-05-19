// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package inboundhttp

import (
	httptypes "github.com/fermyon/spin-go-sdk/internal/fermyon/spin/http-types"
)

// Exports represents the caller-defined exports from "fermyon:spin/inbound-http".
var Exports struct {
	// HandleRequest represents the caller-defined, exported function "handle-request".
	//
	//	handle-request: func(req: request) -> response
	HandleRequest func(req httptypes.Request) (result httptypes.Response)
}

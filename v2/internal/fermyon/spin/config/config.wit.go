// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package config represents the imported interface "fermyon:spin/config".
package config

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Error represents the variant "fermyon:spin/config#error".
//
//	variant error {
//		provider(string),
//		invalid-key(string),
//		invalid-schema(string),
//		other(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorProvider returns a [Error] of case "provider".
func ErrorProvider(data string) Error {
	return cm.New[Error](0, data)
}

// Provider returns a non-nil *[string] if [Error] represents the variant case "provider".
func (self *Error) Provider() *string {
	return cm.Case[string](self, 0)
}

// ErrorInvalidKey returns a [Error] of case "invalid-key".
func ErrorInvalidKey(data string) Error {
	return cm.New[Error](1, data)
}

// InvalidKey returns a non-nil *[string] if [Error] represents the variant case "invalid-key".
func (self *Error) InvalidKey() *string {
	return cm.Case[string](self, 1)
}

// ErrorInvalidSchema returns a [Error] of case "invalid-schema".
func ErrorInvalidSchema(data string) Error {
	return cm.New[Error](2, data)
}

// InvalidSchema returns a non-nil *[string] if [Error] represents the variant case "invalid-schema".
func (self *Error) InvalidSchema() *string {
	return cm.Case[string](self, 2)
}

// ErrorOther returns a [Error] of case "other".
func ErrorOther(data string) Error {
	return cm.New[Error](3, data)
}

// Other returns a non-nil *[string] if [Error] represents the variant case "other".
func (self *Error) Other() *string {
	return cm.Case[string](self, 3)
}

// GetConfig represents the imported function "get-config".
//
// Get a configuration value for the current component.
// The config key must match one defined in in the component manifest.
//
//	get-config: func(key: string) -> result<string, error>
//
//go:nosplit
func GetConfig(key string) (result cm.Result[ErrorShape, string, Error]) {
	key0, key1 := cm.LowerString(key)
	wasmimport_GetConfig((*uint8)(key0), (uint32)(key1), &result)
	return
}

//go:wasmimport fermyon:spin/config get-config
//go:noescape
func wasmimport_GetConfig(key0 *uint8, key1 uint32, result *cm.Result[ErrorShape, string, Error])

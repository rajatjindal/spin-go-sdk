// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package redistypes represents the imported interface "fermyon:spin/redis-types".
package redistypes

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Error represents the enum "fermyon:spin/redis-types#error".
//
// General purpose error.
//
//	enum error {
//		success,
//		error
//	}
type Error uint8

const (
	ErrorSuccess Error = iota
	ErrorError
)

var stringsError = [2]string{
	"success",
	"error",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Error) String() string {
	return stringsError[e]
}

// Payload represents the list "fermyon:spin/redis-types#payload".
//
// The message payload.
//
//	type payload = list<u8>
type Payload cm.List[uint8]

// RedisParameter represents the variant "fermyon:spin/redis-types#redis-parameter".
//
// A parameter type for the general-purpose `execute` function.
//
//	variant redis-parameter {
//		int64(s64),
//		binary(payload),
//	}
type RedisParameter cm.Variant[uint8, Payload, int64]

// RedisParameterInt64 returns a [RedisParameter] of case "int64".
func RedisParameterInt64(data int64) RedisParameter {
	return cm.New[RedisParameter](0, data)
}

// Int64 returns a non-nil *[int64] if [RedisParameter] represents the variant case "int64".
func (self *RedisParameter) Int64() *int64 {
	return cm.Case[int64](self, 0)
}

// RedisParameterBinary returns a [RedisParameter] of case "binary".
func RedisParameterBinary(data Payload) RedisParameter {
	return cm.New[RedisParameter](1, data)
}

// Binary returns a non-nil *[Payload] if [RedisParameter] represents the variant case "binary".
func (self *RedisParameter) Binary() *Payload {
	return cm.Case[Payload](self, 1)
}

// RedisResult represents the variant "fermyon:spin/redis-types#redis-result".
//
// A return type for the general-purpose `execute` function.
//
//	variant redis-result {
//		nil,
//		status(string),
//		int64(s64),
//		binary(payload),
//	}
type RedisResult cm.Variant[uint8, string, int64]

// RedisResultNil returns a [RedisResult] of case "nil".
func RedisResultNil() RedisResult {
	var data struct{}
	return cm.New[RedisResult](0, data)
}

// Nil returns true if [RedisResult] represents the variant case "nil".
func (self *RedisResult) Nil() bool {
	return self.Tag() == 0
}

// RedisResultStatus returns a [RedisResult] of case "status".
func RedisResultStatus(data string) RedisResult {
	return cm.New[RedisResult](1, data)
}

// Status returns a non-nil *[string] if [RedisResult] represents the variant case "status".
func (self *RedisResult) Status() *string {
	return cm.Case[string](self, 1)
}

// RedisResultInt64 returns a [RedisResult] of case "int64".
func RedisResultInt64(data int64) RedisResult {
	return cm.New[RedisResult](2, data)
}

// Int64 returns a non-nil *[int64] if [RedisResult] represents the variant case "int64".
func (self *RedisResult) Int64() *int64 {
	return cm.Case[int64](self, 2)
}

// RedisResultBinary returns a [RedisResult] of case "binary".
func RedisResultBinary(data Payload) RedisResult {
	return cm.New[RedisResult](3, data)
}

// Binary returns a non-nil *[Payload] if [RedisResult] represents the variant case "binary".
func (self *RedisResult) Binary() *Payload {
	return cm.Case[Payload](self, 3)
}

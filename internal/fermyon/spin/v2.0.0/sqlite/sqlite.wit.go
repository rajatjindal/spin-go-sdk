// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package sqlite represents the imported interface "fermyon:spin/sqlite@2.0.0".
package sqlite

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Connection represents the imported resource "fermyon:spin/sqlite@2.0.0#connection".
//
// A handle to an open sqlite instance
//
//	resource connection
type Connection cm.Resource

// ResourceDrop represents the imported resource-drop for resource "connection".
//
// Drops a resource handle.
//
//go:nosplit
func (self Connection) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ConnectionResourceDrop((uint32)(self0))
	return
}

//go:wasmimport fermyon:spin/sqlite@2.0.0 [resource-drop]connection
//go:noescape
func wasmimport_ConnectionResourceDrop(self0 uint32)

// ConnectionOpen represents the imported static function "open".
//
// Open a connection to a named database instance.
//
// If `database` is "default", the default instance is opened.
//
// `error::no-such-database` will be raised if the `name` is not recognized.
//
//	open: static func(database: string) -> result<connection, error>
//
//go:nosplit
func ConnectionOpen(database string) (result cm.Result[ErrorShape, Connection, Error]) {
	database0, database1 := cm.LowerString(database)
	wasmimport_ConnectionOpen((*uint8)(database0), (uint32)(database1), &result)
	return
}

//go:wasmimport fermyon:spin/sqlite@2.0.0 [static]connection.open
//go:noescape
func wasmimport_ConnectionOpen(database0 *uint8, database1 uint32, result *cm.Result[ErrorShape, Connection, Error])

// Execute represents the imported method "execute".
//
// Execute a statement returning back data if there is any
//
//	execute: func(statement: string, parameters: list<value>) -> result<query-result,
//	error>
//
//go:nosplit
func (self Connection) Execute(statement string, parameters cm.List[Value]) (result cm.Result[QueryResultShape, QueryResult, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	statement0, statement1 := cm.LowerString(statement)
	parameters0, parameters1 := cm.LowerList(parameters)
	wasmimport_ConnectionExecute((uint32)(self0), (*uint8)(statement0), (uint32)(statement1), (*Value)(parameters0), (uint32)(parameters1), &result)
	return
}

//go:wasmimport fermyon:spin/sqlite@2.0.0 [method]connection.execute
//go:noescape
func wasmimport_ConnectionExecute(self0 uint32, statement0 *uint8, statement1 uint32, parameters0 *Value, parameters1 uint32, result *cm.Result[QueryResultShape, QueryResult, Error])

// Error represents the variant "fermyon:spin/sqlite@2.0.0#error".
//
// The set of errors which may be raised by functions in this interface
//
//	variant error {
//		no-such-database,
//		access-denied,
//		invalid-connection,
//		database-full,
//		io(string),
//	}
type Error cm.Variant[uint8, cm.StringShape, string]

// ErrorNoSuchDatabase returns a [Error] of case "no-such-database".
//
// The host does not recognize the database name requested.
func ErrorNoSuchDatabase() Error {
	var data struct{}
	return cm.New[Error](0, data)
}

// NoSuchDatabase returns true if [Error] represents the variant case "no-such-database".
func (self *Error) NoSuchDatabase() bool {
	return self.Tag() == 0
}

// ErrorAccessDenied returns a [Error] of case "access-denied".
//
// The requesting component does not have access to the specified database (which
// may or may not exist).
func ErrorAccessDenied() Error {
	var data struct{}
	return cm.New[Error](1, data)
}

// AccessDenied returns true if [Error] represents the variant case "access-denied".
func (self *Error) AccessDenied() bool {
	return self.Tag() == 1
}

// ErrorInvalidConnection returns a [Error] of case "invalid-connection".
//
// The provided connection is not valid
func ErrorInvalidConnection() Error {
	var data struct{}
	return cm.New[Error](2, data)
}

// InvalidConnection returns true if [Error] represents the variant case "invalid-connection".
func (self *Error) InvalidConnection() bool {
	return self.Tag() == 2
}

// ErrorDatabaseFull returns a [Error] of case "database-full".
//
// The database has reached its capacity
func ErrorDatabaseFull() Error {
	var data struct{}
	return cm.New[Error](3, data)
}

// DatabaseFull returns true if [Error] represents the variant case "database-full".
func (self *Error) DatabaseFull() bool {
	return self.Tag() == 3
}

// ErrorIO returns a [Error] of case "io".
//
// Some implementation-specific error has occurred (e.g. I/O)
func ErrorIO(data string) Error {
	return cm.New[Error](4, data)
}

// IO returns a non-nil *[string] if [Error] represents the variant case "io".
func (self *Error) IO() *string {
	return cm.Case[string](self, 4)
}

// Value represents the variant "fermyon:spin/sqlite@2.0.0#value".
//
// A single column's result from a database query
//
//	variant value {
//		integer(s64),
//		real(f64),
//		text(string),
//		blob(list<u8>),
//		null,
//	}
type Value cm.Variant[uint8, cm.StringShape, int64]

// ValueInteger returns a [Value] of case "integer".
func ValueInteger(data int64) Value {
	return cm.New[Value](0, data)
}

// Integer returns a non-nil *[int64] if [Value] represents the variant case "integer".
func (self *Value) Integer() *int64 {
	return cm.Case[int64](self, 0)
}

// ValueReal returns a [Value] of case "real".
func ValueReal(data float64) Value {
	return cm.New[Value](1, data)
}

// Real returns a non-nil *[float64] if [Value] represents the variant case "real".
func (self *Value) Real() *float64 {
	return cm.Case[float64](self, 1)
}

// ValueText returns a [Value] of case "text".
func ValueText(data string) Value {
	return cm.New[Value](2, data)
}

// Text returns a non-nil *[string] if [Value] represents the variant case "text".
func (self *Value) Text() *string {
	return cm.Case[string](self, 2)
}

// ValueBlob returns a [Value] of case "blob".
func ValueBlob(data cm.List[uint8]) Value {
	return cm.New[Value](3, data)
}

// Blob returns a non-nil *[cm.List[uint8]] if [Value] represents the variant case "blob".
func (self *Value) Blob() *cm.List[uint8] {
	return cm.Case[cm.List[uint8]](self, 3)
}

// ValueNull returns a [Value] of case "null".
func ValueNull() Value {
	var data struct{}
	return cm.New[Value](4, data)
}

// Null returns true if [Value] represents the variant case "null".
func (self *Value) Null() bool {
	return self.Tag() == 4
}

// RowResult represents the record "fermyon:spin/sqlite@2.0.0#row-result".
//
// A set of values for each of the columns in a query-result
//
//	record row-result {
//		values: list<value>,
//	}
type RowResult struct {
	Values cm.List[Value]
}

// QueryResult represents the record "fermyon:spin/sqlite@2.0.0#query-result".
//
// A result of a query
//
//	record query-result {
//		columns: list<string>,
//		rows: list<row-result>,
//	}
type QueryResult struct {
	// The names of the columns retrieved in the query
	Columns cm.List[string]

	// the row results each containing the values for all the columns for a given row
	Rows cm.List[RowResult]
}

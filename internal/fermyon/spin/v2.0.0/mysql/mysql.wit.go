// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package mysql represents the imported interface "fermyon:spin/mysql@2.0.0".
package mysql

import (
	rdbmstypes "github.com/fermyon/spin-go-sdk/internal/fermyon/spin/v2.0.0/rdbms-types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// Connection represents the imported resource "fermyon:spin/mysql@2.0.0#connection".
//
// A connection to a MySQL database.
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

//go:wasmimport fermyon:spin/mysql@2.0.0 [resource-drop]connection
//go:noescape
func wasmimport_ConnectionResourceDrop(self0 uint32)

// ConnectionOpen represents the imported static function "open".
//
// Open a connection to the MySQL instance at `address`.
//
//	open: static func(address: string) -> result<connection, error>
//
//go:nosplit
func ConnectionOpen(address string) (result cm.Result[ErrorShape, Connection, rdbmstypes.Error]) {
	address0, address1 := cm.LowerString(address)
	wasmimport_ConnectionOpen((*uint8)(address0), (uint32)(address1), &result)
	return
}

//go:wasmimport fermyon:spin/mysql@2.0.0 [static]connection.open
//go:noescape
func wasmimport_ConnectionOpen(address0 *uint8, address1 uint32, result *cm.Result[ErrorShape, Connection, rdbmstypes.Error])

// Execute represents the imported method "execute".
//
// execute command to the database: insert, update, delete
//
//	execute: func(statement: string, params: list<parameter-value>) -> result<_, error>
//
//go:nosplit
func (self Connection) Execute(statement string, params cm.List[rdbmstypes.ParameterValue]) (result cm.Result[ErrorShape, struct{}, rdbmstypes.Error]) {
	self0 := cm.Reinterpret[uint32](self)
	statement0, statement1 := cm.LowerString(statement)
	params0, params1 := cm.LowerList(params)
	wasmimport_ConnectionExecute((uint32)(self0), (*uint8)(statement0), (uint32)(statement1), (*rdbmstypes.ParameterValue)(params0), (uint32)(params1), &result)
	return
}

//go:wasmimport fermyon:spin/mysql@2.0.0 [method]connection.execute
//go:noescape
func wasmimport_ConnectionExecute(self0 uint32, statement0 *uint8, statement1 uint32, params0 *rdbmstypes.ParameterValue, params1 uint32, result *cm.Result[ErrorShape, struct{}, rdbmstypes.Error])

// Query represents the imported method "query".
//
// query the database: select
//
//	query: func(statement: string, params: list<parameter-value>) -> result<row-set,
//	error>
//
//go:nosplit
func (self Connection) Query(statement string, params cm.List[rdbmstypes.ParameterValue]) (result cm.Result[RowSetShape, rdbmstypes.RowSet, rdbmstypes.Error]) {
	self0 := cm.Reinterpret[uint32](self)
	statement0, statement1 := cm.LowerString(statement)
	params0, params1 := cm.LowerList(params)
	wasmimport_ConnectionQuery((uint32)(self0), (*uint8)(statement0), (uint32)(statement1), (*rdbmstypes.ParameterValue)(params0), (uint32)(params1), &result)
	return
}

//go:wasmimport fermyon:spin/mysql@2.0.0 [method]connection.query
//go:noescape
func wasmimport_ConnectionQuery(self0 uint32, statement0 *uint8, statement1 uint32, params0 *rdbmstypes.ParameterValue, params1 uint32, result *cm.Result[RowSetShape, rdbmstypes.RowSet, rdbmstypes.Error])

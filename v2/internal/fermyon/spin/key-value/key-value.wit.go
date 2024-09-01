// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package keyvalue represents the imported interface "fermyon:spin/key-value".
package keyvalue

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Store represents the u32 "fermyon:spin/key-value#store".
//
// A handle to an open key-value store
//
//	type store = u32
type Store uint32

// Error represents the variant "fermyon:spin/key-value#error".
//
// The set of errors which may be raised by functions in this interface
//
//	variant error {
//		store-table-full,
//		no-such-store,
//		access-denied,
//		invalid-store,
//		no-such-key,
//		io(string),
//	}
type Error cm.Variant[uint8, string, string]

// ErrorStoreTableFull returns a [Error] of case "store-table-full".
//
// Too many stores have been opened simultaneously. Closing one or more
// stores prior to retrying may address this.
func ErrorStoreTableFull() Error {
	var data struct{}
	return cm.New[Error](0, data)
}

// StoreTableFull returns true if [Error] represents the variant case "store-table-full".
func (self *Error) StoreTableFull() bool {
	return self.Tag() == 0
}

// ErrorNoSuchStore returns a [Error] of case "no-such-store".
//
// The host does not recognize the store name requested.  Defining and
// configuring a store with that name in a runtime configuration file
// may address this.
func ErrorNoSuchStore() Error {
	var data struct{}
	return cm.New[Error](1, data)
}

// NoSuchStore returns true if [Error] represents the variant case "no-such-store".
func (self *Error) NoSuchStore() bool {
	return self.Tag() == 1
}

// ErrorAccessDenied returns a [Error] of case "access-denied".
//
// The requesting component does not have access to the specified store
// (which may or may not exist).
func ErrorAccessDenied() Error {
	var data struct{}
	return cm.New[Error](2, data)
}

// AccessDenied returns true if [Error] represents the variant case "access-denied".
func (self *Error) AccessDenied() bool {
	return self.Tag() == 2
}

// ErrorInvalidStore returns a [Error] of case "invalid-store".
//
// The store handle provided is not recognized, i.e. it was either never
// opened or has been closed.
func ErrorInvalidStore() Error {
	var data struct{}
	return cm.New[Error](3, data)
}

// InvalidStore returns true if [Error] represents the variant case "invalid-store".
func (self *Error) InvalidStore() bool {
	return self.Tag() == 3
}

// ErrorNoSuchKey returns a [Error] of case "no-such-key".
//
// No key-value tuple exists for the specified key in the specified
// store.
func ErrorNoSuchKey() Error {
	var data struct{}
	return cm.New[Error](4, data)
}

// NoSuchKey returns true if [Error] represents the variant case "no-such-key".
func (self *Error) NoSuchKey() bool {
	return self.Tag() == 4
}

// ErrorIO returns a [Error] of case "io".
//
// Some implementation-specific error has occurred (e.g. I/O)
func ErrorIO(data string) Error {
	return cm.New[Error](5, data)
}

// IO returns a non-nil *[string] if [Error] represents the variant case "io".
func (self *Error) IO() *string {
	return cm.Case[string](self, 5)
}

// Open represents the imported function "open".
//
// Open the store with the specified name.
//
// If `name` is "default", the default store is opened.  Otherwise,
// `name` must refer to a store defined and configured in a runtime
// configuration file supplied with the application.
//
// `error::no-such-store` will be raised if the `name` is not recognized.
//
//	open: func(name: string) -> result<store, error>
//
//go:nosplit
func Open(name string) (result cm.Result[ErrorShape, Store, Error]) {
	name0, name1 := cm.LowerString(name)
	wasmimport_Open((*uint8)(name0), (uint32)(name1), &result)
	return
}

//go:wasmimport fermyon:spin/key-value open
//go:noescape
func wasmimport_Open(name0 *uint8, name1 uint32, result *cm.Result[ErrorShape, Store, Error])

// Get represents the imported function "get".
//
// Get the value associated with the specified `key` from the specified
// `store`.
//
// `error::invalid-store` will be raised if `store` is not a valid handle
// to an open store, and `error::no-such-key` will be raised if there is no
// tuple for `key` in `store`.
//
//	get: func(store: store, key: string) -> result<list<u8>, error>
//
//go:nosplit
func Get(store Store, key string) (result cm.Result[ErrorShape, cm.List[uint8], Error]) {
	store0 := (uint32)(store)
	key0, key1 := cm.LowerString(key)
	wasmimport_Get((uint32)(store0), (*uint8)(key0), (uint32)(key1), &result)
	return
}

//go:wasmimport fermyon:spin/key-value get
//go:noescape
func wasmimport_Get(store0 uint32, key0 *uint8, key1 uint32, result *cm.Result[ErrorShape, cm.List[uint8], Error])

// Set represents the imported function "set".
//
// Set the `value` associated with the specified `key` in the specified
// `store`, overwriting any existing value.
//
// `error::invalid-store` will be raised if `store` is not a valid handle
// to an open store.
//
//	set: func(store: store, key: string, value: list<u8>) -> result<_, error>
//
//go:nosplit
func Set(store Store, key string, value cm.List[uint8]) (result cm.Result[Error, struct{}, Error]) {
	store0 := (uint32)(store)
	key0, key1 := cm.LowerString(key)
	value0, value1 := cm.LowerList(value)
	wasmimport_Set((uint32)(store0), (*uint8)(key0), (uint32)(key1), (*uint8)(value0), (uint32)(value1), &result)
	return
}

//go:wasmimport fermyon:spin/key-value set
//go:noescape
func wasmimport_Set(store0 uint32, key0 *uint8, key1 uint32, value0 *uint8, value1 uint32, result *cm.Result[Error, struct{}, Error])

// Delete represents the imported function "delete".
//
// Delete the tuple with the specified `key` from the specified `store`.
//
// `error::invalid-store` will be raised if `store` is not a valid handle
// to an open store.  No error is raised if a tuple did not previously
// exist for `key`.
//
//	delete: func(store: store, key: string) -> result<_, error>
//
//go:nosplit
func Delete(store Store, key string) (result cm.Result[Error, struct{}, Error]) {
	store0 := (uint32)(store)
	key0, key1 := cm.LowerString(key)
	wasmimport_Delete((uint32)(store0), (*uint8)(key0), (uint32)(key1), &result)
	return
}

//go:wasmimport fermyon:spin/key-value delete
//go:noescape
func wasmimport_Delete(store0 uint32, key0 *uint8, key1 uint32, result *cm.Result[Error, struct{}, Error])

// Exists represents the imported function "exists".
//
// Return whether a tuple exists for the specified `key` in the specified
// `store`.
//
// `error::invalid-store` will be raised if `store` is not a valid handle
// to an open store.
//
//	exists: func(store: store, key: string) -> result<bool, error>
//
//go:nosplit
func Exists(store Store, key string) (result cm.Result[ErrorShape, bool, Error]) {
	store0 := (uint32)(store)
	key0, key1 := cm.LowerString(key)
	wasmimport_Exists((uint32)(store0), (*uint8)(key0), (uint32)(key1), &result)
	return
}

//go:wasmimport fermyon:spin/key-value exists
//go:noescape
func wasmimport_Exists(store0 uint32, key0 *uint8, key1 uint32, result *cm.Result[ErrorShape, bool, Error])

// GetKeys represents the imported function "get-keys".
//
// Return a list of all the keys in the specified `store`.
//
// `error::invalid-store` will be raised if `store` is not a valid handle
// to an open store.
//
//	get-keys: func(store: store) -> result<list<string>, error>
//
//go:nosplit
func GetKeys(store Store) (result cm.Result[ErrorShape, cm.List[string], Error]) {
	store0 := (uint32)(store)
	wasmimport_GetKeys((uint32)(store0), &result)
	return
}

//go:wasmimport fermyon:spin/key-value get-keys
//go:noescape
func wasmimport_GetKeys(store0 uint32, result *cm.Result[ErrorShape, cm.List[string], Error])

// Close represents the imported function "close".
//
// Close the specified `store`.
//
// This has no effect if `store` is not a valid handle to an open store.
//
//	close: func(store: store)
//
//go:nosplit
func Close(store Store) {
	store0 := (uint32)(store)
	wasmimport_Close((uint32)(store0))
	return
}

//go:wasmimport fermyon:spin/key-value close
//go:noescape
func wasmimport_Close(store0 uint32)

// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package rdbmstypes represents the imported interface "fermyon:spin/rdbms-types".
package rdbmstypes

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// DbDataType represents the enum "fermyon:spin/rdbms-types#db-data-type".
//
//	enum db-data-type {
//		boolean,
//		int8,
//		int16,
//		int32,
//		int64,
//		uint8,
//		uint16,
//		uint32,
//		uint64,
//		floating32,
//		floating64,
//		str,
//		binary,
//		other
//	}
type DbDataType uint8

const (
	DbDataTypeBoolean DbDataType = iota
	DbDataTypeInt8
	DbDataTypeInt16
	DbDataTypeInt32
	DbDataTypeInt64
	DbDataTypeUint8
	DbDataTypeUint16
	DbDataTypeUint32
	DbDataTypeUint64
	DbDataTypeFloating32
	DbDataTypeFloating64
	DbDataTypeStr
	DbDataTypeBinary
	DbDataTypeOther
)

var stringsDbDataType = [14]string{
	"boolean",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"floating32",
	"floating64",
	"str",
	"binary",
	"other",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DbDataType) String() string {
	return stringsDbDataType[e]
}

// DbValue represents the variant "fermyon:spin/rdbms-types#db-value".
//
//	variant db-value {
//		boolean(bool),
//		int8(s8),
//		int16(s16),
//		int32(s32),
//		int64(s64),
//		uint8(u8),
//		uint16(u16),
//		uint32(u32),
//		uint64(u64),
//		floating32(f32),
//		floating64(f64),
//		str(string),
//		binary(list<u8>),
//		db-null,
//		unsupported,
//	}
type DbValue cm.Variant[uint8, string, int64]

// DbValueBoolean returns a [DbValue] of case "boolean".
func DbValueBoolean(data bool) DbValue {
	return cm.New[DbValue](0, data)
}

// Boolean returns a non-nil *[bool] if [DbValue] represents the variant case "boolean".
func (self *DbValue) Boolean() *bool {
	return cm.Case[bool](self, 0)
}

// DbValueInt8 returns a [DbValue] of case "int8".
func DbValueInt8(data int8) DbValue {
	return cm.New[DbValue](1, data)
}

// Int8 returns a non-nil *[int8] if [DbValue] represents the variant case "int8".
func (self *DbValue) Int8() *int8 {
	return cm.Case[int8](self, 1)
}

// DbValueInt16 returns a [DbValue] of case "int16".
func DbValueInt16(data int16) DbValue {
	return cm.New[DbValue](2, data)
}

// Int16 returns a non-nil *[int16] if [DbValue] represents the variant case "int16".
func (self *DbValue) Int16() *int16 {
	return cm.Case[int16](self, 2)
}

// DbValueInt32 returns a [DbValue] of case "int32".
func DbValueInt32(data int32) DbValue {
	return cm.New[DbValue](3, data)
}

// Int32 returns a non-nil *[int32] if [DbValue] represents the variant case "int32".
func (self *DbValue) Int32() *int32 {
	return cm.Case[int32](self, 3)
}

// DbValueInt64 returns a [DbValue] of case "int64".
func DbValueInt64(data int64) DbValue {
	return cm.New[DbValue](4, data)
}

// Int64 returns a non-nil *[int64] if [DbValue] represents the variant case "int64".
func (self *DbValue) Int64() *int64 {
	return cm.Case[int64](self, 4)
}

// DbValueUint8 returns a [DbValue] of case "uint8".
func DbValueUint8(data uint8) DbValue {
	return cm.New[DbValue](5, data)
}

// Uint8 returns a non-nil *[uint8] if [DbValue] represents the variant case "uint8".
func (self *DbValue) Uint8() *uint8 {
	return cm.Case[uint8](self, 5)
}

// DbValueUint16 returns a [DbValue] of case "uint16".
func DbValueUint16(data uint16) DbValue {
	return cm.New[DbValue](6, data)
}

// Uint16 returns a non-nil *[uint16] if [DbValue] represents the variant case "uint16".
func (self *DbValue) Uint16() *uint16 {
	return cm.Case[uint16](self, 6)
}

// DbValueUint32 returns a [DbValue] of case "uint32".
func DbValueUint32(data uint32) DbValue {
	return cm.New[DbValue](7, data)
}

// Uint32 returns a non-nil *[uint32] if [DbValue] represents the variant case "uint32".
func (self *DbValue) Uint32() *uint32 {
	return cm.Case[uint32](self, 7)
}

// DbValueUint64 returns a [DbValue] of case "uint64".
func DbValueUint64(data uint64) DbValue {
	return cm.New[DbValue](8, data)
}

// Uint64 returns a non-nil *[uint64] if [DbValue] represents the variant case "uint64".
func (self *DbValue) Uint64() *uint64 {
	return cm.Case[uint64](self, 8)
}

// DbValueFloating32 returns a [DbValue] of case "floating32".
func DbValueFloating32(data float32) DbValue {
	return cm.New[DbValue](9, data)
}

// Floating32 returns a non-nil *[float32] if [DbValue] represents the variant case "floating32".
func (self *DbValue) Floating32() *float32 {
	return cm.Case[float32](self, 9)
}

// DbValueFloating64 returns a [DbValue] of case "floating64".
func DbValueFloating64(data float64) DbValue {
	return cm.New[DbValue](10, data)
}

// Floating64 returns a non-nil *[float64] if [DbValue] represents the variant case "floating64".
func (self *DbValue) Floating64() *float64 {
	return cm.Case[float64](self, 10)
}

// DbValueStr returns a [DbValue] of case "str".
func DbValueStr(data string) DbValue {
	return cm.New[DbValue](11, data)
}

// Str returns a non-nil *[string] if [DbValue] represents the variant case "str".
func (self *DbValue) Str() *string {
	return cm.Case[string](self, 11)
}

// DbValueBinary returns a [DbValue] of case "binary".
func DbValueBinary(data cm.List[uint8]) DbValue {
	return cm.New[DbValue](12, data)
}

// Binary returns a non-nil *[cm.List[uint8]] if [DbValue] represents the variant case "binary".
func (self *DbValue) Binary() *cm.List[uint8] {
	return cm.Case[cm.List[uint8]](self, 12)
}

// DbValueDbNull returns a [DbValue] of case "db-null".
func DbValueDbNull() DbValue {
	var data struct{}
	return cm.New[DbValue](13, data)
}

// DbNull returns true if [DbValue] represents the variant case "db-null".
func (self *DbValue) DbNull() bool {
	return self.Tag() == 13
}

// DbValueUnsupported returns a [DbValue] of case "unsupported".
func DbValueUnsupported() DbValue {
	var data struct{}
	return cm.New[DbValue](14, data)
}

// Unsupported returns true if [DbValue] represents the variant case "unsupported".
func (self *DbValue) Unsupported() bool {
	return self.Tag() == 14
}

// ParameterValue represents the variant "fermyon:spin/rdbms-types#parameter-value".
//
//	variant parameter-value {
//		boolean(bool),
//		int8(s8),
//		int16(s16),
//		int32(s32),
//		int64(s64),
//		uint8(u8),
//		uint16(u16),
//		uint32(u32),
//		uint64(u64),
//		floating32(f32),
//		floating64(f64),
//		str(string),
//		binary(list<u8>),
//		db-null,
//	}
type ParameterValue cm.Variant[uint8, string, int64]

// ParameterValueBoolean returns a [ParameterValue] of case "boolean".
func ParameterValueBoolean(data bool) ParameterValue {
	return cm.New[ParameterValue](0, data)
}

// Boolean returns a non-nil *[bool] if [ParameterValue] represents the variant case "boolean".
func (self *ParameterValue) Boolean() *bool {
	return cm.Case[bool](self, 0)
}

// ParameterValueInt8 returns a [ParameterValue] of case "int8".
func ParameterValueInt8(data int8) ParameterValue {
	return cm.New[ParameterValue](1, data)
}

// Int8 returns a non-nil *[int8] if [ParameterValue] represents the variant case "int8".
func (self *ParameterValue) Int8() *int8 {
	return cm.Case[int8](self, 1)
}

// ParameterValueInt16 returns a [ParameterValue] of case "int16".
func ParameterValueInt16(data int16) ParameterValue {
	return cm.New[ParameterValue](2, data)
}

// Int16 returns a non-nil *[int16] if [ParameterValue] represents the variant case "int16".
func (self *ParameterValue) Int16() *int16 {
	return cm.Case[int16](self, 2)
}

// ParameterValueInt32 returns a [ParameterValue] of case "int32".
func ParameterValueInt32(data int32) ParameterValue {
	return cm.New[ParameterValue](3, data)
}

// Int32 returns a non-nil *[int32] if [ParameterValue] represents the variant case "int32".
func (self *ParameterValue) Int32() *int32 {
	return cm.Case[int32](self, 3)
}

// ParameterValueInt64 returns a [ParameterValue] of case "int64".
func ParameterValueInt64(data int64) ParameterValue {
	return cm.New[ParameterValue](4, data)
}

// Int64 returns a non-nil *[int64] if [ParameterValue] represents the variant case "int64".
func (self *ParameterValue) Int64() *int64 {
	return cm.Case[int64](self, 4)
}

// ParameterValueUint8 returns a [ParameterValue] of case "uint8".
func ParameterValueUint8(data uint8) ParameterValue {
	return cm.New[ParameterValue](5, data)
}

// Uint8 returns a non-nil *[uint8] if [ParameterValue] represents the variant case "uint8".
func (self *ParameterValue) Uint8() *uint8 {
	return cm.Case[uint8](self, 5)
}

// ParameterValueUint16 returns a [ParameterValue] of case "uint16".
func ParameterValueUint16(data uint16) ParameterValue {
	return cm.New[ParameterValue](6, data)
}

// Uint16 returns a non-nil *[uint16] if [ParameterValue] represents the variant case "uint16".
func (self *ParameterValue) Uint16() *uint16 {
	return cm.Case[uint16](self, 6)
}

// ParameterValueUint32 returns a [ParameterValue] of case "uint32".
func ParameterValueUint32(data uint32) ParameterValue {
	return cm.New[ParameterValue](7, data)
}

// Uint32 returns a non-nil *[uint32] if [ParameterValue] represents the variant case "uint32".
func (self *ParameterValue) Uint32() *uint32 {
	return cm.Case[uint32](self, 7)
}

// ParameterValueUint64 returns a [ParameterValue] of case "uint64".
func ParameterValueUint64(data uint64) ParameterValue {
	return cm.New[ParameterValue](8, data)
}

// Uint64 returns a non-nil *[uint64] if [ParameterValue] represents the variant case "uint64".
func (self *ParameterValue) Uint64() *uint64 {
	return cm.Case[uint64](self, 8)
}

// ParameterValueFloating32 returns a [ParameterValue] of case "floating32".
func ParameterValueFloating32(data float32) ParameterValue {
	return cm.New[ParameterValue](9, data)
}

// Floating32 returns a non-nil *[float32] if [ParameterValue] represents the variant case "floating32".
func (self *ParameterValue) Floating32() *float32 {
	return cm.Case[float32](self, 9)
}

// ParameterValueFloating64 returns a [ParameterValue] of case "floating64".
func ParameterValueFloating64(data float64) ParameterValue {
	return cm.New[ParameterValue](10, data)
}

// Floating64 returns a non-nil *[float64] if [ParameterValue] represents the variant case "floating64".
func (self *ParameterValue) Floating64() *float64 {
	return cm.Case[float64](self, 10)
}

// ParameterValueStr returns a [ParameterValue] of case "str".
func ParameterValueStr(data string) ParameterValue {
	return cm.New[ParameterValue](11, data)
}

// Str returns a non-nil *[string] if [ParameterValue] represents the variant case "str".
func (self *ParameterValue) Str() *string {
	return cm.Case[string](self, 11)
}

// ParameterValueBinary returns a [ParameterValue] of case "binary".
func ParameterValueBinary(data cm.List[uint8]) ParameterValue {
	return cm.New[ParameterValue](12, data)
}

// Binary returns a non-nil *[cm.List[uint8]] if [ParameterValue] represents the variant case "binary".
func (self *ParameterValue) Binary() *cm.List[uint8] {
	return cm.Case[cm.List[uint8]](self, 12)
}

// ParameterValueDbNull returns a [ParameterValue] of case "db-null".
func ParameterValueDbNull() ParameterValue {
	var data struct{}
	return cm.New[ParameterValue](13, data)
}

// DbNull returns true if [ParameterValue] represents the variant case "db-null".
func (self *ParameterValue) DbNull() bool {
	return self.Tag() == 13
}

// Column represents the record "fermyon:spin/rdbms-types#column".
//
//	record column {
//		name: string,
//		data-type: db-data-type,
//	}
type Column struct {
	Name     string
	DataType DbDataType
}

// Row represents the list "fermyon:spin/rdbms-types#row".
//
//	type row = list<db-value>
type Row cm.List[DbValue]

// RowSet represents the record "fermyon:spin/rdbms-types#row-set".
//
//	record row-set {
//		columns: list<column>,
//		rows: list<row>,
//	}
type RowSet struct {
	Columns cm.List[Column]
	Rows    cm.List[Row]
}

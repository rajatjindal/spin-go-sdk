// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package ipnamelookup represents the imported interface "wasi:sockets/ip-name-lookup@0.2.0".
package ipnamelookup

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/fermyon/spin-go-sdk/v2/internal/wasi/io/v0.2.0/poll"
	"github.com/fermyon/spin-go-sdk/v2/internal/wasi/sockets/v0.2.0/network"
)

// ResolveAddressStream represents the imported resource "wasi:sockets/ip-name-lookup@0.2.0#resolve-address-stream".
//
//	resource resolve-address-stream
type ResolveAddressStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "resolve-address-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self ResolveAddressStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [resource-drop]resolve-address-stream
//go:noescape
func wasmimport_ResolveAddressStreamResourceDrop(self0 uint32)

// ResolveNextAddress represents the imported method "resolve-next-address".
//
// Returns the next address from the resolver.
//
// This function should be called multiple times. On each call, it will
// return the next address in connection order preference. If all
// addresses have been exhausted, this function returns `none`.
//
// This function never returns IPv4-mapped IPv6 addresses.
//
// # Typical errors
// - `name-unresolvable`:          Name does not exist or has no suitable associated
// IP addresses. (EAI_NONAME, EAI_NODATA, EAI_ADDRFAMILY)
// - `temporary-resolver-failure`: A temporary failure in name resolution occurred.
// (EAI_AGAIN)
// - `permanent-resolver-failure`: A permanent failure in name resolution occurred.
// (EAI_FAIL)
// - `would-block`:                A result is not available yet. (EWOULDBLOCK, EAGAIN)
//
//	resolve-next-address: func() -> result<option<ip-address>, error-code>
//
//go:nosplit
func (self ResolveAddressStream) ResolveNextAddress() (result cm.Result[OptionIPAddressShape, cm.Option[network.IPAddress], network.ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResolveNextAddress((uint32)(self0), &result)
	return
}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [method]resolve-address-stream.resolve-next-address
//go:noescape
func wasmimport_ResolveAddressStreamResolveNextAddress(self0 uint32, result *cm.Result[OptionIPAddressShape, cm.Option[network.IPAddress], network.ErrorCode])

// Subscribe represents the imported method "subscribe".
//
// Create a `pollable` which will resolve once the stream is ready for I/O.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self ResolveAddressStream) Subscribe() (result poll.Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_ResolveAddressStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[poll.Pollable]((uint32)(result0))
	return
}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 [method]resolve-address-stream.subscribe
//go:noescape
func wasmimport_ResolveAddressStreamSubscribe(self0 uint32) (result0 uint32)

// ResolveAddresses represents the imported function "resolve-addresses".
//
// Resolve an internet host name to a list of IP addresses.
//
// Unicode domain names are automatically converted to ASCII using IDNA encoding.
// If the input is an IP address string, the address is parsed and returned
// as-is without making any external requests.
//
// See the wasi-socket proposal README.md for a comparison with getaddrinfo.
//
// This function never blocks. It either immediately fails or immediately
// returns successfully with a `resolve-address-stream` that can be used
// to (asynchronously) fetch the results.
//
// # Typical errors
// - `invalid-argument`: `name` is a syntactically invalid domain name or IP address.
//
// # References:
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getaddrinfo.html>
// - <https://man7.org/linux/man-pages/man3/getaddrinfo.3.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/ws2tcpip/nf-ws2tcpip-getaddrinfo>
// - <https://man.freebsd.org/cgi/man.cgi?query=getaddrinfo&sektion=3>
//
//	resolve-addresses: func(network: borrow<network>, name: string) -> result<resolve-address-stream,
//	error-code>
//
//go:nosplit
func ResolveAddresses(network_ network.Network, name string) (result cm.Result[ResolveAddressStream, ResolveAddressStream, network.ErrorCode]) {
	network0 := cm.Reinterpret[uint32](network_)
	name0, name1 := cm.LowerString(name)
	wasmimport_ResolveAddresses((uint32)(network0), (*uint8)(name0), (uint32)(name1), &result)
	return
}

//go:wasmimport wasi:sockets/ip-name-lookup@0.2.0 resolve-addresses
//go:noescape
func wasmimport_ResolveAddresses(network0 uint32, name0 *uint8, name1 uint32, result *cm.Result[ResolveAddressStream, ResolveAddressStream, network.ErrorCode])

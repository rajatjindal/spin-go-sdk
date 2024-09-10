module examples/http-outbound/hello

go 1.22.6

require github.com/fermyon/spin-go-sdk/v2 v2.0.0

require (
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/ydnar/wasm-tools-go v0.1.5 // indirect
)

replace github.com/fermyon/spin-go-sdk/v2 => ../../../

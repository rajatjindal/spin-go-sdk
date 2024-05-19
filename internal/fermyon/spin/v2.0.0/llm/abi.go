// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

package llm

import (
	"github.com/ydnar/wasm-tools-go/cm"
	"unsafe"
)

// InferencingResultShape is used for storage in variant or result types.
type InferencingResultShape struct {
	shape [unsafe.Sizeof(InferencingResult{})]byte
}

func lower_InferencingParams(v InferencingParams) (f0 uint32, f1 float32, f2 uint32, f3 float32, f4 uint32, f5 float32) {
	f0 = (uint32)(v.MaxTokens)
	f1 = (float32)(v.RepeatPenalty)
	f2 = (uint32)(v.RepeatPenaltyLastNTokenCount)
	f3 = (float32)(v.Temperature)
	f4 = (uint32)(v.TopK)
	f5 = (float32)(v.TopP)
	return
}

func lower_OptionInferencingParams(v cm.Option[InferencingParams]) (f0 uint32, f1 uint32, f2 float32, f3 uint32, f4 float32, f5 uint32, f6 float32) {
	some := v.Some()
	if some != nil {
		f0 = 1
		v1, v2, v3, v4, v5, v6 := lower_InferencingParams(*some)
		f1 = (uint32)(v1)
		f2 = (float32)(v2)
		f3 = (uint32)(v3)
		f4 = (float32)(v4)
		f5 = (uint32)(v5)
		f6 = (float32)(v6)
	}
	return
}

// EmbeddingsResultShape is used for storage in variant or result types.
type EmbeddingsResultShape struct {
	shape [unsafe.Sizeof(EmbeddingsResult{})]byte
}

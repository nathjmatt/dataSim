package packer

import (
	"encoding/binary"
	"math"
)

const (
	float64_size_in_bytes = 8
)

//	Converts a slice of float64 into a slice of bytes.
//
// samples: the slice of float64s to convert to a byte slice
//
// bigEndian: stores the float64s in BigEndian when true, LittleEndian when false
func PackFloat64(samples []float64, bigEndian bool) []byte {

	byteSamples := make([]byte, len(samples)*float64_size_in_bytes)

	index := 0
	for i := 0; i < len(samples); i++ {
		if bigEndian {
			binary.BigEndian.PutUint64(byteSamples[index:], math.Float64bits(samples[i]))
		} else {
			binary.LittleEndian.PutUint64(byteSamples, math.Float64bits(samples[i]))
		}
		index += float64_size_in_bytes
	}
	return byteSamples
}

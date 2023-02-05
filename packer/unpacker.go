package packer

import (
	"bytes"
	"encoding/binary"
)

//	Unpacks a slice of bytes as a slice of float64
//
// data: the slice of bytes to convert to a slice of float64s
//
// bigEndian: true if the data is stored as BigEndian, false if stored as LittleEndian
func UnpackFloat64(data []byte, bigEndian bool) []float64 {

	dataLen := len(data) / 8
	p := bytes.NewBuffer(data)

	samples := make([]float64, dataLen)

	if bigEndian {
		binary.Read(p, binary.BigEndian, samples)
	} else {
		binary.Read(p, binary.LittleEndian, samples)
	}

	return samples
}

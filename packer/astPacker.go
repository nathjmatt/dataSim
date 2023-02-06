package packer

import (
	"encoding/binary"
	"time"
)

type TwoCyclePacket struct {
	messageLength uint32
	messageType   uint32
	messageCount  uint32

	ia [64]int32
	ib [64]int32
	ic [64]int32
	va [64]int32
	vb [64]int32
	vc [64]int32

	pia [8]uint32
	pib [8]uint32
	pic [8]uint32
	pva [8]uint32
	pvb [8]uint32
	pvc [8]uint32

	// p3mag stuff

	timestamp uint64

	// hif relay word rows

	sdiA    int32
	sdiB    int32
	sdiC    int32
	sdiRefA int32
	sdiRefB int32
	sdiRefC int32
}

const (
	two_cycle_packet_length_in_bytes = 2604
	two_cycle_packet_type            = 2
)

func (d *InstantRawData) SetIa(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.ia); i++ {
		d.ia[i] = samples[i]
	}
}
func (d *InstantRawData) SetIb(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.ib); i++ {
		d.ib[i] = samples[i]
	}
}
func (d *InstantRawData) SetIc(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.ic); i++ {
		d.ic[i] = samples[i]
	}
}
func (d *InstantRawData) SetVa(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.va); i++ {
		d.va[i] = samples[i]
	}
}
func (d *InstantRawData) SetVb(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.vb); i++ {
		d.vb[i] = samples[i]
	}
}
func (d *InstantRawData) SetVc(samples []float64) {
	for i := 0; i < len(samples) || i < len(d.vc); i++ {
		d.vc[i] = samples[i]
	}
}

type InstantRawData struct {
	ia [64]float64
	ib [64]float64
	ic [64]float64
	va [64]float64
	vb [64]float64
	vc [64]float64
}

func NewTwoCyclePacket(chData *InstantRawData) *TwoCyclePacket {
	frame := &TwoCyclePacket{
		messageLength: two_cycle_packet_length_in_bytes,
		messageType:   two_cycle_packet_type,
	}

	for i := 0; i < 64; i++ {
		frame.ia[i] = int32(chData.ia[i] * 1000)
		frame.ib[i] = int32(chData.ib[i] * 1000)
		frame.ic[i] = int32(chData.ic[i] * 1000)
		frame.va[i] = int32(chData.va[i] * 100)
		frame.vb[i] = int32(chData.vb[i] * 100)
		frame.vc[i] = int32(chData.vc[i] * 100)
	}

	frame.timestamp = uint64(time.Now().UnixMilli())

	return frame
}

func (packet *TwoCyclePacket) ToBytes() []byte {
	b := make([]byte, two_cycle_packet_length_in_bytes)

	index := 0
	binary.BigEndian.PutUint32(b, packet.messageLength)
	index += 4
	binary.BigEndian.PutUint32(b[index:], packet.messageType)
	index += 4
	binary.BigEndian.PutUint32(b[index:], packet.messageCount)
	index += 4

	// for i := 0; i < 16; i++ {
	// 	fmt.Println(packet.ia[i])
	// }

	// Pack the current channels
	for i := 0; i < len(packet.ia); i++ {
		binary.BigEndian.PutUint32(b[index:], uint32(packet.ia[i]))
		index += 4
		binary.BigEndian.PutUint32(b[index:], uint32(packet.ib[i]))
		index += 4
		binary.BigEndian.PutUint32(b[index:], uint32(packet.ic[i]))
		index += 4

	}

	// Pack the voltage channels
	for i := 0; i < len(packet.va); i++ {
		binary.BigEndian.PutUint32(b[index:], uint32(packet.va[i]))
		index += 4
		binary.BigEndian.PutUint32(b[index:], uint32(packet.vb[i]))
		index += 4
		binary.BigEndian.PutUint32(b[index:], uint32(packet.vc[i]))
		index += 4

	}

	// Pack the phasor current channels
	for i := 0; i < len(packet.pia); i++ {
		binary.BigEndian.PutUint32(b[index:], packet.pia[i])
		index += 4
		binary.BigEndian.PutUint32(b[index:], packet.pib[i])
		index += 4
		binary.BigEndian.PutUint32(b[index:], packet.pic[i])
		index += 4
	}

	// Pack the phasor voltage channels
	for i := 0; i < len(packet.pva); i++ {
		binary.BigEndian.PutUint32(b[index:], packet.pva[i])
		index += 4
		binary.BigEndian.PutUint32(b[index:], packet.pvb[i])
		index += 4
		binary.BigEndian.PutUint32(b[index:], packet.pvc[i])
		index += 4
	}

	// p3 stuff
	for i := 0; i < 8; i++ {

	}
	// p angle stuff
	for i := 0; i < 8; i++ {

	}

	index = 2188
	binary.BigEndian.PutUint64(b[index:], packet.timestamp)
	index += 8

	// hif relay word bits

	index = 2580
	// sdi values

	return b
}

package nuls

import (
	"encoding/binary"
)

type Na uint64

type VarInt uint64

func (v VarInt) sizeOf() int {
	if v < 0 {
		// 1 marker + 8 data bytes
		return 9
	}
	if v < 253 {
		// 1 data byte
		return 1
	}
	if v <= 0xFFFF {
		// 1 marker + 2 data bytes
		return 3
	}
	if v <= 0xFFFFFFFF {
		// 1 marker + 4 data bytes
		return 5
	}
	return 9
}

func (v VarInt) encode() []byte {
	switch v.sizeOf() {
	case 1:
		return []byte{byte(v)}
	case 3:
		return []byte{byte(253), byte(v), byte(v >> 8)}
	case 5:
		bytes:=make([]byte,5)
		bytes[0]=byte(254)
		binary.LittleEndian.PutUint32(bytes[1:],uint32(v))
		return bytes
	default:
		bytes:=make([]byte,9)
		bytes[0]=byte(255)
		binary.LittleEndian.PutUint64(bytes[1:],uint64(v))
		return bytes

	}
}

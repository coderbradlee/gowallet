package nuls

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"errors"
)

const TX_TYPE_TRANSFER = 2;

type Transaction struct {
	BaseData
	Types       int16  `json:type`
	time        uint64 `json:time`
	ScriptSig   []byte `json:scriptSig`
	txData      Serializable
	coinData    CoinData
	remark      []byte `json:remark`
	hash        NulsDigestData
	blockHeight int64
}

func (t *Transaction) Size() int {
	size := 0
	size += sizeOfUint16()
	size += sizeOfUint48()
	size += sizeOfBytes(t.remark)
	size += sizeOfNulsData(t.txData)
	size += sizeOfNulsData(&t.coinData)
	size += sizeOfBytes(t.ScriptSig)
	return size
}
func (d *Transaction) Parse(b []byte) error {
	err := d.BaseData.Parse(b)
	if err != nil {
		return err
	}
	return nil
}

func (d *Transaction) SerializeToByte() ([]byte, error) {
	size := d.Size()
	if size == 0 {
		return PlaceHolder, nil
	}
	var bf bytes.Buffer
	temp := make([]byte, 8)

	binary.LittleEndian.PutUint16(temp[0:2], uint16(d.Types))
	bf.Write(temp[0:2])

	writeUint48(d.time, &bf)

	writeBytesWithLength(d.remark, &bf)

	writeNulsData(d.txData, &bf)
	writeNulsData(&d.coinData, &bf)
	writeBytesWithLength(d.ScriptSig, &bf)

	if len(bf.Bytes()) != size {
		return nil, errors.New(fmt.Sprintf("data serialize error %T", d))
	}
	return bf.Bytes(), nil
}
func (d *Transaction) SerializeForHash() []byte {
	size := d.Size() - sizeOfBytes(d.ScriptSig)

	var bf bytes.Buffer
	if size == 0 {
		bf.Write(PlaceHolder)
	} else {
		writeVarInt(VarInt(d.Types), &bf)
		writeVarInt(VarInt(d.time), &bf)
		writeBytesWithLength(d.remark, &bf)
		writeNulsData(d.txData, &bf)
		writeNulsData(&d.coinData, &bf)
	}
	return bf.Bytes()
}
func (d *Transaction) setHash(b NulsDigestData) {
	d.hash = b
}

func (d *Transaction) getHash() NulsDigestData {
	return d.hash
}

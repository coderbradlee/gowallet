package nuls

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"errors"
)

type Coin struct {
	BaseData
	owner    []byte
	na       Na
	lockTime uint64
	from     *Coin
}

func (d *Coin) Size() int {
	size := 0
	size += sizeOfBytes(d.owner)
	size += sizeOfInt64()
	size += sizeOfUint48()
	return size
}

func (d *Coin) Parse(b []byte) error {
	err := d.BaseData.Parse(b)
	if err != nil {
		return err
	}
	return nil
}

func (d *Coin) SerializeToByte() ([]byte, error) {
	size := d.Size()
	if size == 0 {
		return PlaceHolder, nil
	}
	var bf bytes.Buffer
	writeBytesWithLength(d.owner, &bf)
	temp := make([]byte, 8)
	binary.LittleEndian.PutUint64(temp, uint64(d.na))
	bf.Write(temp)
	writeUint48(d.lockTime, &bf)

	if len(bf.Bytes()) != size {
		return nil, errors.New(fmt.Sprintf("data serialize error %T", d))
	}
	return bf.Bytes(), nil
}

package nuls

import (
	"bytes"
	"encoding/hex"
)

const(
	SIGN_ALG_ECC byte=0
	SIGN_ALG_DEFAULT byte=SIGN_ALG_ECC
)

type NulsSignData struct {
	BaseData
	signAlgType byte
	signBytes []byte
}

func (d *NulsSignData) Size() int {
	size:=sizeOfBytes(d.signBytes)+1
	return size
}

func (d *NulsSignData) Parse(b []byte) error {
	err := d.BaseData.Parse(b)
	if err != nil {
		return err
	}
	return nil
}

func (d *NulsSignData) SerializeToByte() ([]byte, error) {
	var bf bytes.Buffer
	bf.Write([]byte{d.signAlgType})
	writeBytesWithLength(d.signBytes,&bf)
	return bf.Bytes(),nil
}

func (d *NulsSignData) String() string {
	b,_:=d.SerializeToByte()
	return hex.EncodeToString(b)
}


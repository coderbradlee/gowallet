package nuls

import (
	"bytes"
)

const DEFAULT_SERIALIZE_LENGTH  = 110
type P2PKHScriptSig struct {
	BaseData
	publicKey []byte
	signData NulsSignData

}

func (d *P2PKHScriptSig) Size() int {
	size:=1+ len(d.publicKey)
	size+=sizeOfNulsData(&d.signData)
	return size
}

func (d *P2PKHScriptSig) Parse(b []byte) error {
	err := d.BaseData.Parse(b)
	if err != nil {
		return err
	}
	return nil
}

func (d *P2PKHScriptSig) SerializeToByte() ([]byte, error) {
	var bf bytes.Buffer
	bf.WriteByte(byte(len(d.publicKey)))
	bf.Write(d.publicKey)
	writeNulsData(&d.signData,&bf)
	return bf.Bytes(),nil
}



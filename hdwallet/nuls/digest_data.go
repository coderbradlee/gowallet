package nuls

import (
	"bytes"
	"errors"
	"fmt"
	"encoding/hex"
	"crypto/sha256"
	"crypto"
)

const (
	HASH_LENGTH=34
	DIGEST_ALG_SHA256=0
	DIGEST_ALG_SHA160=1
)

type NulsDigestData struct {
	BaseData
	digestBytes  []byte
	digestAlgType byte
}

func (d *NulsDigestData) SerializeToByte() ([]byte,error) {
	var bf bytes.Buffer
	bf.WriteByte(byte(d.digestAlgType))
	writeBytesWithLength(d.digestBytes,&bf)
	return bf.Bytes(),nil
}

func (d *NulsDigestData) Size() int {
	return sizeOfBytes(d.digestBytes)+1
}

func (d *NulsDigestData) Parse(b []byte) error {
	if b==nil || len(b)==0 || (len(b)==4 && bytes.Equal(b, PlaceHolder)) {
		return errors.New(fmt.Sprintf("data parse error %T",d))
	}
	return nil
}

func (d *NulsDigestData) GetDigestHex() string  {
	b,_:=d.SerializeToByte()
	return hex.EncodeToString(b)
}

func calcDigestData(b []byte,digestAlgType byte) *NulsDigestData  {
	if digestAlgType==DIGEST_ALG_SHA256{
		c:=sha256.Sum256(b)
		hash:=sha256.Sum256(c[:])
		return &NulsDigestData{digestAlgType:0,digestBytes:hash[:]}
	}else if digestAlgType==DIGEST_ALG_SHA160{
		c:=sha256.Sum256(b)
		digest := crypto.RIPEMD160.New()
		digest.Write(c[:])
		r:=digest.Sum(nil)
		return &NulsDigestData{digestAlgType:1,digestBytes:r}
	}
	return nil
}
package nuls

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	PlaceHolder=[]byte{0xFF,0xFF,0xFF,0xFF}
)
type Serializable interface {
	SerializeToByte() ([]byte,error)
	Parse([]byte) error
	Size() int
}

type BaseData struct {

}

func (d *BaseData) SerializeToByte() ([]byte,error) {
	size:=d.Size()
	if size==0{
		return PlaceHolder,nil
	}
	return nil,nil
	/*var b bytes.Buffer
	d.serializeToBuffer(&b)
	if len(b.Bytes())!=size{
		return nil,errors.New(fmt.Sprintf("data serialize error %T",d))
	}
	return b.Bytes(),nil*/
}

func (d *BaseData) Size() int {
	return 0
}

func (d *BaseData) Parse(b []byte) error {
	if b==nil || len(b)==0 || (len(b)==4 && bytes.Equal(b, PlaceHolder)) {
		return errors.New(fmt.Sprintf("data parse error %T",d))
	}
	return nil
}


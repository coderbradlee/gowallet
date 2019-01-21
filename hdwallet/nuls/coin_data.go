package nuls

import (
	"bytes"
)

type CoinData struct {
	BaseData
	from []Coin
	to []Coin

}

func (d *CoinData) Size() int {
	size:=0
	if d.from!=nil{
		size=sizeOfVarInt(uint64(len(d.from)))
		for _,v:=range d.from{
			size+=sizeOfNulsData(&v)
		}
	}else{
		size=sizeOfVarInt(uint64(0))
	}
	if d.to!=nil{
		size+=sizeOfVarInt(uint64(len(d.to)))
		for _,v:=range d.to{
			size+=sizeOfNulsData(&v)
		}
	}else{
		size+=sizeOfVarInt(uint64(0))
	}
	return size
}

func (d *CoinData) Parse(b []byte) error {
	err := d.BaseData.Parse(b)
	if err != nil {
		return err
	}
	return nil
}

func (d *CoinData) SerializeToByte() ([]byte, error) {
	var bf bytes.Buffer
	if d.from!=nil{
		bf.Write(VarInt(len(d.from)).encode())
		for _,v:=range d.from{
			writeNulsData(&v,&bf)
		}
	}else{
		bf.Write(VarInt(0).encode())
	}
	if d.to!=nil{
		bf.Write(VarInt(len(d.to)).encode())
		for _,v:=range d.to{
			writeNulsData(&v,&bf)
		}
	}else{
		bf.Write(VarInt(0).encode())
	}
	return bf.Bytes(),nil
}

func (d *CoinData) AddTo(coin Coin)  {
	d.to=append(d.to,coin)
}

func (d *CoinData) AddFrom(coin Coin)  {
	d.from=append(d.from,coin)
}

package nuls

import (
	"bytes"
	"encoding/binary"
)

func sizeOfUint16() int  {
	return 2
}
func sizeOfInt16() int  {
	return 2
}

func sizeOfUint32() int  {
	return 4
}

func sizeOfInt32() int  {
	return 4
}

func sizeOfUint48() int  {
	return 6
}
func sizeOfInt64() int  {
	return 8
}

func sizeOfVarInt(v uint64) int  {
	return VarInt(v).sizeOf()
}

func sizeOfBytes(b []byte) int  {
	if b==nil{
		return 1
	}
	return VarInt(len(b)).sizeOf()+ len(b)
}
func sizeOfNulsData(data Serializable) int  {
	if data==nil {
		return len(PlaceHolder)
	}
	len:=data.Size()
	if len==0{
		return 1
	} else{
		return len
	}
}

func writeBytesWithLength(bytes []byte,bf *bytes.Buffer)  {
	if bytes ==nil || len(bytes)==0{
		bf.Write(VarInt(0).encode())
	}else{
		bf.Write(VarInt(len(bytes)).encode())
		bf.Write(bytes)
	}
}

func writeUint48(time uint64,bf *bytes.Buffer)  {
	temp:=make([]byte,8)
	binary.LittleEndian.PutUint64(temp,time)
	bf.Write(temp[0:6])
}

func writeNulsData(data Serializable,bf *bytes.Buffer)  {
	if data==nil{
		bf.Write(PlaceHolder)
	}else {
		ds,_:=data.SerializeToByte()
		bf.Write(ds)
	}
}

func writeVarInt(val VarInt,bf *bytes.Buffer)  {
	bf.Write(val.encode())
}
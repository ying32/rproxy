package main

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestHttp(t *testing.T) {
	//
	enVal := EncodeVerify()
	t.Log(enVal)
	t.Log(DecodeVerify(enVal))

	head := make([]byte, 5)
	head[0] = 1
	head[1] = 2
	head[2] = 3
	head[3] = 4
	head[4] = 5
	t.Log(head[0:3])
	t.Log(head[1:3])
	t.Log(head[3:])

	pHead := TPackageHead{}
	pHead.Head = PacketHead
	pHead.Version = PacketVersion
	pHead.Cmd = PacketCmd1
	pHead.DataLen = 111
	raw := bytes.NewBuffer([]byte{})
	binary.Write(raw, binary.LittleEndian, pHead)
	///
	phead2 := TPackageHead{}
	raw2 := bytes.NewBuffer(raw.Bytes())
	binary.Read(raw2, binary.LittleEndian, &phead2)
	t.Log(phead2)

}

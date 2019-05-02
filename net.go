package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

const (
	PacketHead uint8 = 0x02
	PacketTail uint8 = 0x03

	PacketCmd1   uint16 = 1000
	PacketVerify uint16 = 1001
	PackageError uint16 = 1002

	PacketVersion uint16 = 0x01
)

type TPackageHead struct {
	Version uint16 // 2
	Cmd     uint16 // 2
	DataLen uint32 // 4
	//Data []byte
	//Tail uint8
}

var (
	// 封包头的结构长度
	PackageHeadLen = 8

	// 验证包长度，因为懒，所以直接固定值
	PackageVerifyLen = 1 + PackageHeadLen + len(verifyVal) + 1
)

/*

  http.ReadRequest()
  http.ReadResponse()
  httputil.DumpRequest()
  httputil.DumpResponse()
*/

// 编码数据
func EncodeCmd(cmd uint16, data []byte) []byte {
	raw := bytes.NewBuffer([]byte{})
	head := TPackageHead{}
	head.Version = PacketVersion
	head.Cmd = cmd
	head.DataLen = uint32(len(data))

	binary.Write(raw, binary.LittleEndian, PacketHead)
	binary.Write(raw, binary.LittleEndian, head)
	raw.Write(data)
	binary.Write(raw, binary.LittleEndian, PacketTail)
	return raw.Bytes()
}

// DecodeHead
func DecodeHead(data []byte) *TPackageHead {
	head := new(TPackageHead)
	raw := bytes.NewBuffer(data)
	binary.Read(raw, binary.LittleEndian, head)

	return head
}

// 编码验证包，为固定长度， 1 + 2 + 2 + 20 + 4 + 1 = 30byte
func EncodeVerify() []byte {
	return EncodeCmd(PacketVerify, verifyVal[:])
}

// 检测数据包，主要对包头，包尾，版本，命令效验，并返回剩余的buff和对应的错误
func checkPackage(data []byte, cmd uint16) (*bytes.Buffer, error) {
	if len(data) <= 0 {
		return nil, errors.New("封包数据长度不正确。")
	}
	// 校验包的首尾
	if !(data[0] == PacketHead && data[len(data)-1] == PacketTail) {
		return nil, errors.New("封包数据首尾不正确。")
	}
	// buff，移除首尾
	raw := bytes.NewBuffer(data[1 : len(data)-1])
	var val uint16
	binary.Read(raw, binary.LittleEndian, &val)
	if val != PacketVersion {
		return nil, errors.New("封包数据版本与服务端不一致。")
	}
	binary.Read(raw, binary.LittleEndian, &val)
	if val != cmd {
		return nil, errors.New("封包数据版命令与当前要求命令不一致。")
	}
	return raw, nil
}

// 解码验证包，并验证是否正则
func DecodeVerify(data []byte) error {
	raw, err := checkPackage(data, PacketVerify)
	if err != nil {
		return err
	}
	var dataLen uint32
	binary.Read(raw, binary.LittleEndian, &dataLen)
	if int(dataLen) != len(verifyVal) {
		return errors.New("数据长度不符合。")
	}
	val := make([]byte, 20)
	raw.Read(val)
	if bytes.Compare(val, verifyVal[:]) != 0 {
		return errors.New("首次连接校验证失败。")
	}
	return nil
}

// 将request 的处理
func EncodeRequest(r *http.Request) ([]byte, error) {
	reqBytes, err := httputil.DumpRequest(r, true)
	if err != nil {
		return nil, err
	}
	// 判断是否为http或者https的标识1字节
	//binary.Write(raw, binary.LittleEndian, bool(r.URL.Scheme == "https"))
	return EncodeCmd(PacketCmd1, reqBytes), err
}

// 将字节转为request
func DecodeRequest(data []byte, port int, isHttps bool) (*http.Request, error) {
	if len(data) <= 100 {
		return nil, errors.New("待解码的字节长度太小")
	}
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	req.Host = "127.0.0.1"
	if port != 80 {
		req.Host += ":" + strconv.Itoa(port)
	}
	scheme := "http"
	if isHttps {
		scheme = "https"
	}
	req.URL, _ = url.Parse(fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI))
	req.RequestURI = ""

	return req, nil
}

// 将response转为字节
func EncodeResponse(r *http.Response) ([]byte, error) {
	respBytes, err := httputil.DumpResponse(r, true)
	if err != nil {
		return nil, err
	}
	return EncodeCmd(PacketCmd1, respBytes), err
}

//// 将字节转为response
func DecodeResponse(data []byte) (*http.Response, error) {

	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(data)), nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//-------------------------一些简化---------------------------
// IO错误检测
func chkIOError(err error) bool {
	return err == io.EOF || err == io.ErrUnexpectedEOF
}

// 写错误
func wError(conn net.Conn, err error) {
	conn.Write(EncodeCmd(PackageError, []byte(err.Error())))
}

// 校验包头，并返回head和相应的错误
func chkHead(conn net.Conn) (*TPackageHead, error) {
	// 检测
	headData := make([]byte, PackageHeadLen)
	n, err := conn.Read(headData)
	// 读取长度不对，或者出错，则继续
	if n != PackageHeadLen {
		return nil, errors.New("读取头长度错误。")
	}
	if err != nil {
		return nil, err
	}
	head, err := DecodeHead(headData)
	if err != nil {
		return nil, err
	}
	log.Println("head:", head)

	// 读取正确：解析数据
	// 版本解析
	if head.Version != PacketVersion {
		return nil, errors.New("包版本与服务器不一致。")
	}
	return head, nil
}

// 校验包尾
func chkTail(conn net.Conn) error {
	tail := make([]byte, 1)
	_, err := conn.Read(tail)
	if err != nil {
		return err
	}
	log.Println("读取尾部数据成功")
	if tail[0] != PacketTail {
		return errors.New("封包尾校验失败。")
	}
	log.Println("尾部数成校验成功。")
	return nil
}

// 读主体
func readBody(conn net.Conn, dataLen uint32) ([]byte, error) {
	// 单页最大不能超过4M
	if dataLen <= 0 || dataLen > 1024*1024*4 {
		return nil, errors.New("数据长度太大。")
	}
	log.Println("数据长度范围正确。")
	buff := make([]byte, dataLen)
	n, err := conn.Read(buff)
	if n != int(dataLen) {
		return nil, errors.New("读取数据度长不对。")
	}
	log.Println("主体长度：", dataLen, "，已经读取：", n)
	if err != nil {
		return nil, err
	}
	log.Println("读取数据正常")

	return buff, nil
}

// 读数据包
func readPackage(conn net.Conn, fn func(cmd uint16, data []byte)) error {
	byteFlag := make([]byte, 1)
	_, err := conn.Read(byteFlag)
	if err != nil {
		return err
	}
	// 检测包头，必须等于这个才
	if byteFlag[0] == PacketHead {
		headBuff := make([]byte, PackageHeadLen)
		_, err := conn.Read(headBuff)
		if err != nil {
			return err
		}
		head := DecodeHead(headBuff)
		if head.Version == PacketVersion {
			if head.DataLen > 0 || head.DataLen <= 1024*1024*4 {
				bodyData := make([]byte, head.DataLen)
				_, err := conn.Read(bodyData)
				if err != nil {
					return err
				}
				_, err = conn.Read(byteFlag)
				if err != nil {
					return err
				}
				if byteFlag[0] == PacketTail {
					fn(head.Cmd, bodyData)
				} else {
					log.Println("包尾不正确")
				}
			} else {
				log.Println("数据太长")
			}
		} else {
			log.Println("版本不一致")
		}
	}
	return nil
}

package librp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

const (
	PacketHead uint8 = 0x02
	PacketTail uint8 = 0x03

	PacketCmd1   uint16 = 1000
	PacketVerify uint16 = 1001
	PackageError uint16 = 1002

	PacketVersion uint16 = 0x01
)

type TPacketHead struct {
	//Head uint8
	Version uint16 // 2
	Cmd     uint16 // 2
	DataLen uint32 // 4
	//Data []byte
	//Tail uint8
}

var (
	// 封包头的结构长度
	PacketHeadLen = 8
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
	head := TPacketHead{}
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
func DecodeHead(data []byte) *TPacketHead {
	head := new(TPacketHead)
	raw := bytes.NewBuffer(data)
	binary.Read(raw, binary.LittleEndian, head)

	return head
}

// 编码验证包，为固定长度， 1 + 2 + 2 + 20 + 4 + 1 = 30byte
func EncodeVerify() []byte {
	return EncodeCmd(PacketVerify, verifyVal[:])
}

// 验证成功回写
func EncodeVerifyOK() []byte {
	return EncodeCmd(PacketVerify, []byte("ok"))
}

// 验证失败回写
func EncodeVerifyFailed() []byte {
	return EncodeCmd(PacketVerify, []byte("failed"))
}

// 将request 的处理
func EncodeRequest(r *http.Request) ([]byte, error) {
	reqBytes, err := httputil.DumpRequest(r, true)
	if err != nil {
		return nil, err
	}
	return EncodeCmd(PacketCmd1, reqBytes), err
}

// 将字节转为request
func DecodeRequest(data []byte, port int, isHttps bool) (*http.Request, error) {
	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	req.Host = "127.0.0.1"
	scheme := "http"
	if isHttps {
		scheme = "https"
		if port != 443 {
			req.Host += ":" + strconv.Itoa(port)
		}
	} else {
		if port != 80 {
			req.Host += ":" + strconv.Itoa(port)
		}
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

// 写错误
func wError(conn net.Conn, err error) error {
	return wData(conn, EncodeCmd(PackageError, []byte(err.Error())))
}

// 写数据
func wData(conn net.Conn, data []byte) error {
	if conn == nil {
		return errors.New("连接无效。")
	}
	// 先不理写入的
	_, err := conn.Write(data)
	return err
}

// 读数据
func rData(conn net.Conn, bLen int) ([]byte, error) {
	bsBuff := bytes.NewBuffer([]byte{})
	bufLen := bLen
	for {
		if bsBuff.Len() >= bLen {
			break
		}
		buf := make([]byte, bufLen)
		nr, err := conn.Read(buf)
		if err != nil {
			return nil, err
		}
		bsBuff.Write(buf[:nr])
		if nr == bLen {
			break
		}
		bufLen = bLen - bsBuff.Len()
	}
	return bsBuff.Bytes(), nil
}

// 读数据包
func readPacket(conn net.Conn, fn func(cmd uint16, data []byte) error) error {
	if conn == nil {
		return errors.New("连接无效。")
	}
	byteFlag := make([]byte, 1)
	_, err := conn.Read(byteFlag)
	if err != nil {
		return err
	}
	// 检测包头，必须等于这个才
	if byteFlag[0] == PacketHead {
		headBuff := make([]byte, PacketHeadLen)
		_, err := conn.Read(headBuff)
		if err != nil {
			return err
		}
		head := DecodeHead(headBuff)
		if head.Version == PacketVersion {
			if head.DataLen > 0 || head.DataLen <= 1024*1024*4 {
				bodyData, err := rData(conn, int(head.DataLen))
				if err != nil {
					return err
				}
				_, err = conn.Read(byteFlag)
				if err != nil {
					return err
				}
				if byteFlag[0] == PacketTail {
					return fn(head.Cmd, bodyData)
				} else {
					Log.E("包尾不正确")
				}
			} else {
				Log.E("数据太长")
			}
		} else {
			Log.E("版本不一致")
		}
	}
	return nil
}

func keepALive(conn net.Conn) {
	if conn.(*net.TCPConn).SetKeepAlive(true) == nil {
		conn.(*net.TCPConn).SetKeepAlivePeriod(time.Duration(1 * time.Minute))
	}
}

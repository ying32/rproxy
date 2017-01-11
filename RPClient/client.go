package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

var (
	serverAddr    = flag.String("addr", "", "连接服务器的地址，如：127.0.0.1:82")
	httpProtocol  = flag.String("protocol", "http", "服务器的协议，可选为http、https")
	httpPort      = flag.Int("port", 0, "要求服务端开启转发的端口，只作参考，不强制要求")
	localhttpPort = flag.Int("localport", 0, "要转发至本地的http端口，不填或者错误端口号则跟httpPort一致！")
)

func main() {
	flag.Parse()
	if *httpProtocol != "http" && *httpProtocol != "https" {
		fmt.Println("http协议错误！")
		return
	}
	if *serverAddr == "" {
		fmt.Println("请输入服务器的地址！")
		return
	}
	if *localhttpPort <= 0 || *localhttpPort >= 65536 {
		fmt.Println("端口配置范围错误！")
		return
	}

	fmt.Println("信息：本次要求服务器开启：", *httpPort, "号端口，将被转发至本地：", *localhttpPort, "端口，使用协议为：", *httpProtocol)
	conn, err := net.Dial("tcp", *serverAddr)
	if err != nil {
		fmt.Println("连接服务端失败：", err.Error())
		return
	}
	defer conn.Close()
	StartClient(conn)
}

func StartClient(conn net.Conn) {
	fmt.Println("已成功连接服务端！地址：", *serverAddr)
	val := make([]byte, 4)
	for {
		_, err := conn.Read(val)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
			break
		}
		if string(val) == "sign" {
			c, err := conn.Read(val)
			if err == nil && c == 4 {
				bodylen := binary.LittleEndian.Uint32(val)
				fmt.Println("服务端发送数据长为：", bodylen)
				if bodylen > 0 {
					bs := make([]byte, bodylen)
					c, err := conn.Read(bs)
					if err == nil && c == int(bodylen) {
						if err = httpMethod(bs, conn); err != nil {
							fmt.Println("httpMethod请求错误：", err)
							conn.Write([]byte("msg0"))
						}
					} else {
						fmt.Println("读取数据长错误:", err)
					}
				}
			}
		} else if string(val) == "msg0" {
			fmt.Println("服务器配置错误，消息懒得读了")
			return
		} else {
			fmt.Println("读取签名错误！")
		}
	}
}

// 请求的方法
func httpMethod(data []byte, conn net.Conn) error {
	if len(data) == 0 {
		return errors.New("无数据可用")
	}
	reader := bytes.NewReader(data)
	// 读取Method
	var nlen int32 = 0
	binary.Read(reader, binary.LittleEndian, &nlen)
	if nlen == 0 || nlen > 10 {
		return errors.New("Method数据长度错误！")
	}
	mbytes := make([]byte, nlen)
	binary.Read(reader, binary.LittleEndian, &mbytes)
	method := string(mbytes)

	// 读取请求的url
	nlen = 0
	binary.Read(reader, binary.LittleEndian, &nlen)
	if nlen == 0 || nlen > 1024 {
		return errors.New("URL数据长度错误！")
	}
	mbytes = make([]byte, nlen)
	binary.Read(reader, binary.LittleEndian, &mbytes)
	url := string(mbytes)
	// 读取body的数据
	var contentlen int64
	binary.Read(reader, binary.LittleEndian, &contentlen)
	body := bytes.NewBuffer([]byte{})
	if contentlen > 0 {
		mbytes = make([]byte, contentlen)
		binary.Read(reader, binary.LittleEndian, mbytes)
		binary.Write(body, binary.LittleEndian, mbytes)
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s://127.0.0.1:%d%s", *httpProtocol, *localhttpPort, url), body)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	retbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sendio := bytes.NewBuffer([]byte{})
	// 写签名
	binary.Write(sendio, binary.LittleEndian, []byte("sign"))
	// 写长度
	binary.Write(sendio, binary.LittleEndian, int32(len(retbytes)))
	// 写数据
	binary.Write(sendio, binary.LittleEndian, retbytes)

	if n, err := conn.Write(sendio.Bytes()); n == sendio.Len() && err == nil {
		fmt.Println("回复成功！")
		return nil
	} else {
		return err
	}
}

package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	tcpPort  = flag.Int("port", 0, "TCP监听的端口")
	httpPort = flag.Int("httpport", 0, "监听的http端口")
)

var (
	clientConn *net.TCPConn
)

func main() {
	flag.Parse()
	if *tcpPort <= 0 || *httpPort >= 65536 {
		fmt.Println("请输入一个正确的TCP端口")
		return
	}
	if *httpPort <= 0 || *httpPort >= 65536 {
		fmt.Println("请输入一个正确的HTTP端口")
	}
	fmt.Println("监听服务端监端口：", *tcpPort)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: *tcpPort, Zone: ""})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	go startHttpServer(*httpPort)
	fmt.Println("正在进行TCP连接监听...")
	startServer(listen)
}

func startServer(l *net.TCPListener) {
	for {
		conn, err := l.AcceptTCP()
		// 一个服务端只对应一个客户端
		if err != nil || clientConn != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		// 处理新的连接
		fmt.Println("处理新的客户端连接")
		go processTCPConn(conn)
	}
}

func startHttpServer(port int) {
	fmt.Println("开始监听HTTP", port, "端口...")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if clientConn != nil {

			httpBody := bytes.NewBuffer([]byte{})
			// 请求的方法
			mBytes := []byte(r.Method)
			binary.Write(httpBody, binary.LittleEndian, int32(len(mBytes)))
			binary.Write(httpBody, binary.LittleEndian, mBytes)

			// 请求的url
			mBytes = []byte(r.URL.String())
			binary.Write(httpBody, binary.LittleEndian, int32(len(mBytes)))
			binary.Write(httpBody, binary.LittleEndian, mBytes)
			// 如果是IO数据
			binary.Write(httpBody, binary.LittleEndian, r.ContentLength)
			if r.ContentLength > 0 {
				bs, _ := ioutil.ReadAll(r.Body)
				binary.Write(httpBody, binary.LittleEndian, bs)
			}

			sendio := bytes.NewBuffer([]byte{})
			binary.Write(sendio, binary.LittleEndian, []byte("sign"))
			binary.Write(sendio, binary.LittleEndian, int32(httpBody.Len()))
			binary.Write(sendio, binary.LittleEndian, httpBody.Bytes())

			mBytes = sendio.Bytes()
			fmt.Println("输出：", string(mBytes))

			if nlen, err := clientConn.Write(mBytes); len(mBytes) != nlen || err != nil {
				fmt.Println("发送错误：已发送：", nlen, "字节， 错误消息：", err)
			} else {
				fmt.Println("本次发送成功！共：", nlen, "个字节")
			}

			if clientConn != nil {

				val := make([]byte, 4)
				_, err := clientConn.Read(val)
				flags := string(val)
				if err == nil && flags == "sign" {
					c, err := clientConn.Read(val)
					if err == nil && c == 4 {
						bs := bytes.NewBuffer(val)
						var bodylen int32
						binary.Read(bs, binary.LittleEndian, &bodylen)
						if bodylen > 0 {
							bs := make([]byte, bodylen)
							if c, err := clientConn.Read(bs); err == nil && c == int(bodylen) {
								w.Write(bs)
								return
							}
						}
					}
				} else if flags == "msg0" {
					fmt.Println("客户端http请求返回失败！")
				}
			}
		}
		w.Write([]byte(""))
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("http服务端错误，消息：", err)
		return
	}
	fmt.Println("HTTP服务结束。")
}

func processTCPConn(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Duration(10 * time.Second))
	defer conn.Close()
	clientConn = conn
	test := make([]byte, 0)
	for {
		_, err := clientConn.Read(test)
		if err != nil {
			fmt.Println("客户端断开连接！ ")
			break
		}
		time.Sleep(time.Duration(time.Second * 3))
	}
	clientConn = nil
	fmt.Println("客户端已经退出！")
}

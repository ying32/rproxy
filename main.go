package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	tcpPort   = flag.Int("tcpport", 0, "Socket连接或者监听的端口")
	httpPort  = flag.Int("httpport", 0, "当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口")
	rpMode    = flag.String("mode", "client", "启动模式，可选为client、server")
	svrAddr   = flag.String("svraddr", "127.0.0.1", "当mode为client时有效，为连接服务器的地址")
	verifyKey = flag.String("vkey", "", "用作客户端与服务端连接时的校验")
)

func main() {
	flag.Parse()
	if *verifyKey == "" {
		logFatalln("必须输入一个验证的key")
	}
	if *tcpPort <= 0 || *tcpPort >= 65536 {
		logFatalln("请输入正确的tcp端口。")
	}
	if *httpPort <= 0 || *httpPort >= 65536 {
		logFatalln("请输入正确的http端口。")
	}
	if *rpMode != "client" && *rpMode != "server" {
		logFatalln("请输入正确的服务启动模式")
	}
	if *rpMode == "server" && *tcpPort == *httpPort {
		logFatalln("tcp端口与http端口不能为同一个。")
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if *rpMode == "client" {
	retry:
		logPrintln("客户端启动，连接服务器：", *svrAddr, "， 端口：", *tcpPort, "， 并开启http服务端，端口为：", *httpPort)
		cli := NewRPClient(fmt.Sprintf("%s:%d", *svrAddr, *tcpPort), *httpPort)
		if err := cli.Start(); err != nil {
			logPrintln(err)
			// 重连
			logPrintln("5秒后重新连接...")
			time.Sleep(time.Second * 5)
			goto retry
		}
		defer cli.Close()
	} else if *rpMode == "server" {
		logPrintln("服务端启动，监听tcp服务端端口：", *tcpPort, "， http服务端端口：", *httpPort)
		svr := NewRPServer(*tcpPort, *httpPort)
		if err := svr.Start(); err != nil {
			logFatalln(err)
		}
		defer svr.Close()
	}
}

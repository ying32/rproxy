package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	tcpPort  = flag.Int("tcpport", 0, "Socket连接或者监听的端口")
	httpPort = flag.Int("httpport", 0, "当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口")
	rpMode   = flag.String("mode", "client", "启动模式，可选为client、server")
	svrAddr  = flag.String("svraddr", "127.0.0.1", "当mode为client时有效，为连接服务器的地址")
)

func main() {
	flag.Parse()
	if *tcpPort <= 0 || *tcpPort >= 65536 {
		log.Fatalln("请输入正确的tcp端口。")
	}
	if *httpPort <= 0 || *httpPort >= 65536 {
		log.Fatalln("请输入正确的http端口。")
	}
	if *rpMode != "client" && *rpMode != "server" {
		log.Fatalln("请输入正确的服务启动模式")
	}
	if *rpMode == "server" && *tcpPort == *httpPort {
		log.Fatalln("tcp端口与http端口不能为同一个。")
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if *rpMode == "client" {
		log.Println("客户端启动，连接：", *svrAddr, "， 端口：", *tcpPort, "， 并开启http服务端，端口为：", *httpPort)
		cli := NewRPClient(fmt.Sprintf("%s:%d", *svrAddr, *tcpPort), *httpPort)
		if err := cli.Start(); err != nil {
			log.Fatalln(err)
		}
		defer cli.Close()
	} else if *rpMode == "server" {
		log.Println("服务端启动，监听tcp服务端端口：", *tcpPort, "， http服务端端口：", *httpPort)
		svr := NewRPServer(*tcpPort, *httpPort)
		if err := svr.Start(); err != nil {
			log.Fatalln(err)
		}
		defer svr.Close()
	}
}

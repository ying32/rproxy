package main

import (
	"flag"
	"fmt"
	"time"

	. "github.com/ying32/rproxy/librp"
)

var (
	tcpPort     = flag.Int("tcpport", 0, "Socket连接或者监听的端口")
	httpPort    = flag.Int("httpport", 0, "当mode为server时为服务端监听端口，当为mode为client时为转发至本地客户端的端口")
	rpMode      = flag.String("mode", "client", "启动模式，可选为client、server")
	svrAddr     = flag.String("svraddr", "127.0.0.1", "当mode为client时有效，为连接服务器的地址")
	verifyKey   = flag.String("vkey", "", "用作客户端与服务端连接时的校验")
	isHTTPS     = flag.Bool("ishttps", false, "httpPort端口是否只用作HTTPS监听")
	tlsCAFile   = flag.String("tlscafile", "", "当ishttps为true时，所需的CA根证书文件")
	tlsCertFile = flag.String("tlscertfile", "", "当ishttps为true时，所需求的TLS证书文件")
	tlsKeyFile  = flag.String("tlskeyfile", "", "当ishttps为true时，所需求的TLS密匙文件")
)

func main() {
	flag.Parse()
	if *verifyKey == "" {
		Log.EF("必须输入一个验证的key")
	}
	if *tcpPort <= 0 || *tcpPort >= 65536 {
		Log.EF("请输入正确的tcp端口。")
	}
	if *httpPort <= 0 || *httpPort >= 65536 {
		Log.EF("请输入正确的http端口。")
	}
	if *rpMode != "client" && *rpMode != "server" {
		Log.EF("请输入正确的服务启动模式")
	}
	if *rpMode == "server" && *tcpPort == *httpPort {
		Log.EF("tcp端口与http端口不能为同一个。")
	}
	if *isHTTPS && (*tlsCertFile == "" || *tlsKeyFile == "" || *tlsCAFile == "") {
		Log.EF("当为HTTPS时，TLS证书不能为空。")
	}

	if *isHTTPS {
		// 添加CA根证书
		AddRootCert(*tlsCAFile)
		SetTLSCertFile(*tlsCertFile)
		SetTLSKeyFile(*tlsKeyFile)
	}

	InitVerifyKey(*verifyKey)

	if *rpMode == "client" {
	retry:
		Log.I("客户端启动，连接服务器：", *svrAddr, "， 端口：", *tcpPort)
		if *isHTTPS {
			Log.I("转发至HTTP服务为HTTPS")
		}
		Log.I("转发至本地HTTP(S)端口：", *httpPort)
		cli := NewRPClient(fmt.Sprintf("%s:%d", *svrAddr, *tcpPort), *httpPort, *isHTTPS)
		if err := cli.Start(); err != nil {
			Log.E(err)
			// 重连
			Log.I("5秒后重新连接...")
			time.Sleep(time.Second * 5)
			goto retry
		}
		defer cli.Close()
	} else if *rpMode == "server" {
		Log.I("TCP服务端已启动，端口：", *tcpPort)
		if *isHTTPS {
			Log.I("当前HTTP服务为HTTPS")
		}
		Log.I("HTTP(S)服务端已开启，端口：", *httpPort)
		svr := NewRPServer(*tcpPort, *httpPort, *isHTTPS)
		if err := svr.Start(); err != nil {
			Log.EF(err)
		}
		defer svr.Close()
	}
}

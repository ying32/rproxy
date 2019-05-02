package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type TRPServer struct {
	tcpPort  int
	httpPort int
	listener net.Listener
	conn     net.Conn
	sync.RWMutex
}

func NewRPServer(tcpPort, httpPort int) *TRPServer {
	s := new(TRPServer)
	s.tcpPort = tcpPort
	s.httpPort = httpPort
	return s
}

func (s *TRPServer) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.tcpPort))
	if err != nil {
		return err
	}
	go s.httpServer()
	return s.tcpServer()
}

func (s *TRPServer) Close() error {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
	if s.listener != nil {
		err := s.listener.Close()
		s.listener = nil
		return err
	}
	return errors.New("TCP实例未创建！")
}

func (s *TRPServer) tcpServer() error {
	var err error
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.cliProcess(conn)
	}
	return err
}

func badRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (s *TRPServer) httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Lock()
		defer s.Unlock()
		logPrintln(r.RequestURI)
		err := s.write(r)
		if err != nil {
			badRequest(w)
			log.Println(err)
			return
		}
		err = s.read(w)
		if err != nil {
			badRequest(w)
			log.Println(err)
			return
		}
	})
	logFatalln(http.ListenAndServe(fmt.Sprintf(":%d", s.httpPort), nil))
}

func (s *TRPServer) cliProcess(conn net.Conn) error {
	//  客户端没有在连接成功后5秒内发送数据则超时
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	buff := make([]byte, PackageVerifyLen)
	n, err := conn.Read(buff)
	if n != PackageVerifyLen || err != nil {
		logPrintln("客户端读取错误，关闭当前连接。")
		conn.Close()
		return err
	}
	// 清除
	conn.SetReadDeadline(time.Time{})
	// 校验，偷懒，直接来
	if err = DecodeVerify(buff); err != nil {
		logPrintln("当前客户端连接校验错误，关闭此客户端:", conn.RemoteAddr())
		//conn.Write(EncodeCmd(PackageError, []byte("验证失败！")))
		conn.Close()
		return err
	}
	// 检测上次已连接的客户端，尝试断开
	if s.conn != nil {
		logPrintln("服务端已有客户端连接！断开之前的:", conn.RemoteAddr())
		s.conn.Close()
		s.conn = nil
	}
	logPrintln("连接新的客户端：", conn.RemoteAddr())
	conn.(*net.TCPConn).SetKeepAlive(true)
	conn.(*net.TCPConn).SetKeepAlivePeriod(time.Duration(2 * time.Second))
	s.conn = conn
	return nil
}

func (s *TRPServer) write(r *http.Request) error {
	if s.conn != nil {
		reqBytes, err := EncodeRequest(r)
		if err != nil {
			return err
		}
		log.Println("当前服务端总字节数：", len(reqBytes))
		_, err = s.conn.Write(reqBytes)
		if err != nil {
			return err
		}
	} else {
		return errors.New("客户端未连接。")
	}
	return nil
}

func (s *TRPServer) read(w http.ResponseWriter) error {
	if s.conn != nil {

		head, err := chkHead(s.conn)
		if err != nil {
			log.Println("error:", err, head)
			return err
		}

		log.Println(err)
		log.Println("当前命令：", head.Cmd, ", Head:", head)
		// 命令解析
		switch head.Cmd {
		case PacketCmd1:
			bodyData, err := readBody(s.conn, head.DataLen)
			if err != nil {
				return err
			}
			// Decode Response
			resp, err := DecodeResponse(bodyData)
			if err != nil {
				return err
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			for k, v := range resp.Header {
				for _, v2 := range v {
					w.Header().Set(k, v2)
				}
			}
			w.WriteHeader(resp.StatusCode)
			w.Write(bodyBytes)

		default:

		}
	}
	return nil
}

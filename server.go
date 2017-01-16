package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type TRPServer struct {
	tcpPort  int
	httpPort int
	listener *net.TCPListener
	conn     *net.TCPConn
}

func NewRPServer(tcpPort, httpPort int) *TRPServer {
	s := new(TRPServer)
	s.tcpPort = tcpPort
	s.httpPort = httpPort
	return s
}

func (s *TRPServer) Start() error {
	var err error
	s.listener, err = net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("0.0.0.0"), s.tcpPort, ""})
	if err != nil {
		return err
	}
	go s.httpserver()
	return s.tcpserver()
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

func (s *TRPServer) tcpserver() error {
	var err error
	for {
		conn, err := s.listener.AcceptTCP()
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

func (s *TRPServer) httpserver() {
	// google总是要请求这个，这里不要了！
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
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
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", s.httpPort), nil))
}

func (s *TRPServer) cliProcess(conn *net.TCPConn) error {
	// 设置5秒超时
	conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
	vval := make([]byte, 20)
	_, err := conn.Read(vval)
	if err != nil {
		log.Println("客户端读超时。客户端地址为：:", conn.RemoteAddr())
		conn.Close()
		return err
	}
	if bytes.Compare(vval, getverifyval()[:]) != 0 {
		log.Println("当前客户端连接校验错误，关闭此客户端:", conn.RemoteAddr())
		conn.Close()
		return err
	}
	// 清除
	conn.SetReadDeadline(time.Time{})

	if s.conn != nil {
		log.Println("服务端已有客户端连接！断开之前的:", conn.RemoteAddr())
		s.conn.Close()
		s.conn = nil
	}
	log.Println("连接新的客户端：", conn.RemoteAddr())
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Duration(2 * time.Second))
	s.conn = conn
	return nil
}

func (s *TRPServer) write(r *http.Request) error {
	if s.conn != nil {
		raw, err := EncodeRequest(r)
		if err != nil {
			return err
		}
		c, err := s.conn.Write(raw)
		if err != nil {
			return err
		}
		if c != len(raw) {
			return errors.New("写出长度与字节长度不一致。")
		}
	} else {
		return errors.New("客户端未连接。")
	}
	return nil
}

func (s *TRPServer) read(w http.ResponseWriter) error {
	if s.conn != nil {
		val := make([]byte, 4) // flag
		_, err := s.conn.Read(val)
		if err != nil {
			return err
		}
		flags := string(val)
		switch flags {
		case "sign":
			_, err = s.conn.Read(val)
			if err != nil {
				return err
			}
			nlen := binary.LittleEndian.Uint32(val)
			if nlen == 0 {
				return errors.New("读取客户端长度错误。")
			}
			raw := make([]byte, nlen)
			c, err := s.conn.Read(raw)
			if err != nil {
				return err
			}
			if c != int(nlen) {
				return errors.New("读取长度错误。")
			}
			resp, err := DecodeResponse(raw)
			if err != nil {
				return err
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			for k, v := range resp.Header {
				for _, v2 := range v {
					w.Header().Add(k, v2)
				}
			}
			w.WriteHeader(resp.StatusCode)
			w.Write(bodyBytes)
		case "msg0":
			return errors.New("客户端返回错误，但我懒得读了！")
		}
	} else {
		return errors.New("客户端未连接。")
	}
	return nil
}

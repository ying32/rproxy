package librp

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

type TRPServer struct {
	IRPObject
	listener net.Listener
	httpSvr  *http.Server
	conn     net.Conn
	sync.RWMutex
}

func NewRPServer() *TRPServer {
	s := new(TRPServer)
	return s
}

func (s *TRPServer) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", conf.TCPPort))
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
	if s.httpSvr != nil {
		s.httpSvr.Close()
		s.httpSvr = nil
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
		if s.listener == nil {
			return errors.New("监听已关闭。")
		}
		conn, err := s.listener.Accept()
		if err != nil {
			Log.E(err)
			continue
		}
		go s.cliProcess(conn)
	}
	return err
}

func badRequest(w http.ResponseWriter) {
	http.Error(w, "请求错误，错误消息请看控制台信息。", http.StatusBadRequest)
}

type THTTPHandler struct {
	http.Handler
	l     sync.RWMutex
	read  func(w http.ResponseWriter) error
	write func(r *http.Request) error
}

func newHTTPHandler(l sync.RWMutex, read func(w http.ResponseWriter) error, write func(r *http.Request) error) *THTTPHandler {
	h := new(THTTPHandler)
	h.l = l
	h.read = read
	h.write = write
	return h
}

func (h *THTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Lock()
	defer h.l.Unlock()
	Log.I(r.Method + " " + r.RequestURI)
	err := h.write(r)
	if err != nil {
		badRequest(w)
		Log.E(err)
		return
	}
	err = h.read(w)
	if err != nil {
		badRequest(w)
		Log.E(err)
		return
	}
}

func (s *TRPServer) httpServer() {

	s.httpSvr = &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Server.HTTPPort),
		Handler: newHTTPHandler(s.RWMutex, s.read, s.write),
		TLSConfig: &tls.Config{
			ClientCAs: conf.certPool,
			// 这个不开启。。。
			//ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	if !conf.IsHTTPS {
		Log.EF(s.httpSvr.ListenAndServe())
	} else {

		Log.EF(s.httpSvr.ListenAndServeTLS(conf.Server.TLSCertFile, conf.Server.TLSKeyFile))
	}
}

func (s *TRPServer) cliProcess(conn net.Conn) error {
	//  客户端没有在连接成功后5秒内发送数据则超时
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	err := readPacket(conn, func(cmd uint16, data []byte) error {
		conn.SetReadDeadline(time.Time{})
		if cmd == PacketVerify {
			if bytes.Compare(data, conf.verifyVal[:]) != 0 {
				return errors.New("首次连接校验证失败。")
			}
		} else {
			return errors.New("首次请求命令不正确。")
		}
		return nil
	})
	if err != nil {
		Log.W("当前客户端连接校验错误，关闭此客户端。")
		conn.Write(EncodeVerifyFailed())
		conn.Close()
		return err
	}
	// 检测上次已连接的客户端，尝试断开
	if s.conn != nil {
		Log.W("服务端已有客户端连接！断开之前的:", IPStr(conn))
		s.conn.Close()
		s.conn = nil
	}
	if _, err := conn.Write(EncodeVerifyOK()); err != nil {
		return err
	}
	Log.I("连接新的客户端：", IPStr(conn))
	s.conn = conn
	keepALive(s.conn)
	return nil
}

func (s *TRPServer) write(r *http.Request) error {
	if s.conn == nil {
		return errors.New("客户端未连接。")
	}
	reqBytes, err := EncodeRequest(r)
	if err != nil {
		return err
	}
	return wData(s.conn, reqBytes)
}

func (s *TRPServer) read(w http.ResponseWriter) error {
	return readPacket(s.conn, func(cmd uint16, data []byte) error {
		switch cmd {
		case PacketCmd1:
			resp, err := DecodeResponse(data)
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

		case PackageError:
			return errors.New(string(data))
		}

		return nil
	})
}

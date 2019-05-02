package librp

import (
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	disabledRedirect = errors.New("disabled redirect.")
)

type TRPClient struct {
	svrAddr  string
	httpPort int
	conn     net.Conn
}

func NewRPClient(svraddr string, httpPort int) *TRPClient {
	c := new(TRPClient)
	c.svrAddr = svraddr
	c.httpPort = httpPort
	return c
}

func (c *TRPClient) Start() error {
	conn, err := net.Dial("tcp", c.svrAddr)
	if err != nil {
		return err
	}
	conn.(*net.TCPConn).SetKeepAlive(true)
	conn.(*net.TCPConn).SetKeepAlivePeriod(time.Duration(2 * time.Second))
	c.conn = conn
	return c.process()
}

func (c *TRPClient) process() error {
	// 首先请求验证
	v := EncodeVerify()
	if _, err := c.conn.Write(v); err != nil {
		return err
	}

	// 如果服务端没有主动关闭链接则说明已经认证成功
	LogPrintln("已连接服务端。")

	doHTTPClient := func(req *http.Request) ([]byte, error) {
		rawQuery := ""
		if req.URL.RawQuery != "" {
			rawQuery = "?" + req.URL.RawQuery
		}
		LogPrintln(req.Method + "  " + req.URL.Path + rawQuery)
		// 请求本地指定的HTTP服务器
		client := new(http.Client)
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return disabledRedirect
		}
		resp, err := client.Do(req)
		disRedirect := err != nil && strings.Contains(err.Error(), disabledRedirect.Error())
		if err != nil && !disRedirect {
			return nil, err
		}
		defer resp.Body.Close()
		if disRedirect {
			resp.Body = nil
			resp.ContentLength = 0
		}

		respBytes, err := EncodeResponse(resp)
		if err != nil {
			return nil, err
		}

		return respBytes, nil
	}

	// read循环
	for {
		err := readPacket(c.conn, func(cmd uint16, data []byte) error {
			switch cmd {
			case PacketCmd1:
				// Decode请求
				req, err := DecodeRequest(data, c.httpPort, false)
				if err != nil {
					//log.Println("解析请求数据失败：", err)
					return wError(c.conn, err)
				}
				respBytes, err := doHTTPClient(req)
				if err != nil {
					//log.Println("请求本地客户端数据失败：", err)
					return wError(c.conn, err)
				}
				_, err = c.conn.Write(respBytes)
				if err != nil {
					// 写出错了，这里要退出
					return err
				}
			}

			return nil
		})
		// read出错，退出
		if err != nil {
			c.Close()
			return err
		}
	}
	return nil
}

func (c *TRPClient) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return errors.New("TCP实例未创建")
}

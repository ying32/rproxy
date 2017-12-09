package main

import (
	"encoding/binary"
	"errors"
	"log"
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

func (c *TRPClient) werror() {
	c.conn.Write([]byte("msg0"))
}

func (c *TRPClient) process() error {
	if _, err := c.conn.Write(getverifyval()); err != nil {
		return err
	}
	val := make([]byte, 4)
	for {
		_, err := c.conn.Read(val)
		if err != nil {
			log.Println("读数据错误，错误：", err)
			return err
		}
		flags := string(val)
		switch flags {
		case "sign":
			_, err := c.conn.Read(val)
			nlen := binary.LittleEndian.Uint32(val)
			log.Println("收到服务端数据，长度：", nlen)
			if nlen <= 0 {
				log.Println("数据长度错误。")
				c.werror()
				continue
			}
			raw := make([]byte, nlen)
			n, err := c.conn.Read(raw)
			if err != nil {
				return err
			}
			if n != int(nlen) {
				log.Printf("读取服务端数据长度错误，已经读取%dbyte，总长度%d字节\n", n, nlen)
				c.werror()
				continue
			}
			req, err := DecodeRequest(raw, *httpPort)
			if err != nil {
				log.Println("DecodeRequest错误：", err)
				c.werror()
				continue
			}

			rawQuery := ""
			if req.URL.RawQuery != "" {
				rawQuery = "?" + req.URL.RawQuery
			}
			log.Println(req.URL.Path + rawQuery)
			client := new(http.Client)
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return disabledRedirect
			}
			resp, err := client.Do(req)
			disRedirect := err != nil && strings.Contains(err.Error(), disabledRedirect.Error())
			if err != nil && !disRedirect {
				log.Println("请求本地客户端错误：", err)
				c.werror()
				continue
			}
			if !disRedirect {
				defer resp.Body.Close()
			} else {
				resp.Body = nil
				resp.ContentLength = 0
			}
			respBytes, err := EncodeResponse(resp)
			if err != nil {
				log.Println("EncodeResponse错误：", err)
				c.werror()
				continue
			}
			n, err = c.conn.Write(respBytes)
			if err != nil {
				log.Println("发送数据错误，错误：", err)
			}
			if n != len(respBytes) {
				log.Printf("发送数据长度错误，已经发送：%dbyte，总字节长：%dbyte\n", n, len(respBytes))
			} else {
				log.Printf("本次请求成功完成，共发送：%dbyte\n", n)
			}

		case "msg0":
			log.Println("服务端返回错误。")
		default:
			// 不知道啥错误，不输出了
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

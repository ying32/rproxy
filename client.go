package main

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	"net/http"
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
	c.conn = conn
	return c.process()
}

func (c *TRPClient) werror() {
	c.conn.Write([]byte("msg0"))
}

func (c *TRPClient) process() error {
	val := make([]byte, 4)
	for {
		_, err := c.conn.Read(val)
		if err != nil {
			return err
		}
		flags := string(val)
		switch flags {
		case "sign":
			_, err := c.conn.Read(val)
			nlen := binary.LittleEndian.Uint32(val)
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
				log.Println("读取数据长度错误。")
				c.werror()
				continue
			}
			req, err := DecodeRequest(raw, *httpPort)
			if err != nil {
				log.Println(err)
				c.werror()
				continue
			}
			log.Println(req.URL.Path)
			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				c.werror()
				continue
			}
			defer resp.Body.Close()
			respBytes, err := EncodeResponse(resp)
			if err != nil {
				log.Println(err)
				c.werror()
				continue
			}
			n, err = c.conn.Write(respBytes)
			if err != nil {
				return err
			}
			if n != len(respBytes) {
				return errors.New("发送数据长度错误。")
			}
		case "msg0":
			return errors.New("服务端返回错误。")
		default:
			return errors.New("服务端未知错误。")
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

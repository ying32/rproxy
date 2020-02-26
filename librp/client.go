package librp

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type TRPClient struct {
	IRPObject
	conn    net.Conn
	running bool
}

func NewRPClient() *TRPClient {
	c := new(TRPClient)
	return c
}

func (c *TRPClient) Start() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.Client.SvrAddr, conf.TCPPort))
	if err != nil {
		return err
	}
	c.running = true
	c.conn = conn
	return c.process()
}

func (c *TRPClient) process() error {
	// 首先请求验证
	if _, err := c.conn.Write(EncodeVerify()); err != nil {
		return err
	}
	err := readPacket(c.conn, func(cmd uint16, data []byte) error {
		if cmd == PacketVerify {
			if string(data) == "ok" {
				return nil
			}
		}
		return errors.New("验证失败。")
	})
	if err != nil {
		return err
	}

	keepALive(c.conn)
	Log.I("已连接服务端。")

	doHTTPClient := func(req *http.Request) ([]byte, error) {
		rawQuery := ""
		if req.URL.RawQuery != "" {
			rawQuery = "?" + req.URL.RawQuery
		}
		Log.I(req.Method + "  " + req.URL.Path + rawQuery)
		// 请求本地指定的HTTP服务器
		client := new(http.Client)

		if conf.IsHTTPS {
			client.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      conf.certPool,
					Certificates: []tls.Certificate{conf.cliCert},
				},
			}
		}

		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		resp, err := client.Do(req)

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		respBytes, err := EncodeResponse(resp)
		if err != nil {
			return nil, err
		}

		return respBytes, nil
	}

	// read循环
	for c.running {
		err := readPacket(c.conn, func(cmd uint16, data []byte) error {
			switch cmd {
			case PacketCmd1:
				// Decode请求

				req, err := DecodeRequest(data, conf.Client.LocalAddr, conf.Client.HTTPPort, conf.IsHTTPS)
				if err != nil {
					return wError(c.conn, err)
				}
				respBytes, err := doHTTPClient(req)
				if err != nil {
					return wError(c.conn, err)
				}
				_, err = c.conn.Write(respBytes)
				if err != nil {
					// 写出错了，这里要退出
					return err
				}
			case PackageError:
				Log.E(string(data))
			}

			return nil
		})
		// read出错，退出
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *TRPClient) Close() error {
	c.running = false
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return errors.New("TCP实例未创建")
}

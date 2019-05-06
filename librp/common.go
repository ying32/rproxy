package librp

import (
	"crypto/sha1"
	"net"
	"strings"
)

// 验证的值
var verifyVal [20]byte

// 初始验证key
func InitVerifyKey(key string) {
	verifyVal = sha1.Sum([]byte("I AM A KEY:" + key))
}

// 从 xxx.xxx.xxx.xxx:xxx格式中取出ip地址
func IPStr(conn net.Conn) string {
	if conn == nil {
		return ""
	}
	ip := conn.RemoteAddr().String()
	i := strings.LastIndex(ip, ":")
	if i == -1 {
		return ip
	}
	return ip[:i]
}

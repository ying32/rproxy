package librp

import "crypto/sha1"

// 验证的值
var verifyVal [20]byte

// 初始验证key
func InitVerifyKey(key string) {
	verifyVal = sha1.Sum([]byte("I AM A KEY:" + key))
}

package main

import (
	"crypto/sha1"
)

var (
	// 验证的值
	verifyVal [20]byte
)

func init() {
	verifyVal = sha1.Sum([]byte("I AM A KEY:" + *verifyKey))
}

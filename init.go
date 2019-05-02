package main

import (
	"crypto/sha1"
)

func init()  {
	verifyVal= sha1.Sum([]byte("I AM A KEY:" + *verifyKey))
}

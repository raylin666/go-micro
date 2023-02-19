package utils

import "unsafe"

func ByteStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

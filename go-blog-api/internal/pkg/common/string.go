package common

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)

}

func GetRandomInt(l int) string {
	str := "0123456789"
	b := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func StringsContain(dst string, lst ...string) bool {
	for _, s := range lst {
		if s == dst {
			return true
		}
	}
	return false
}
func IntContain(dst int, lst ...int) bool {
	for _, s := range lst {
		if s == dst {
			return true
		}
	}
	return false
}
func StringsContainsAny(source []string, dest ...string) bool {
	for _, b := range dest {
		for _, a := range source {
			if b == a {
				return true
			}
		}
	}
	return false
}

// MarshalUnEscape 不转义Marshal
func MarshalUnEscape(params interface{}) []byte {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(params)
	encodeStr := strings.Trim(bf.String(), "\n")
	return stringTobyteSlice(encodeStr)
}

// stringTobyteSlice 字符串转byte
func stringTobyteSlice(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))
}

package utils

import "fmt"

func GenPassword(password, salt string) string {
	md5, _ := Md5(password)
	result, _ := Md5(fmt.Sprintf("%s%s", md5, salt))
	return result
}

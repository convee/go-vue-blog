package utils

import (
	"os"
)

// 判断所给路径文件/文件夹是否存在

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// RemoveFile 移除文件
func RemoveFile(path string) bool {
	if FileExists(path) {
		err := os.Remove(path)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

// GetPemStringFromFile 获取密钥内容
func GetPemStringFromFile(path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	// 获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	_, _ = file.Read(buf)

	return string(buf), nil
}

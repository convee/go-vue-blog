package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/convee/go-blog-api/configs"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/vansante/go-ffprobe.v2"
	image2 "image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetVideoInfo(path string) (*ffprobe.ProbeData, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelFn()
	ffprobePath := configs.Conf.App.FfprobePath
	ffprobe.SetFFProbeBinPath(ffprobePath)
	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func DownloadFile(uri string) (string, error) {
	if configs.Conf.App.Env == "prod" {
		// todo 可将外网地址转换为内容，使用内网带宽下载素材
	}
	fileName := uuid.NewV4().String() + filepath.Base(uri)
	path := GetTmpFilePath() + "/" + fileName
	// 文件下载超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("素材地址无法访问")
	}
	defer resp.Body.Close()
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return path, nil
}

func GetFileType(format string) string {
	if StringsContain(format, "jpg", "jpeg", "png", "gif") {
		return "image"
	} else if StringsContain(format, "avi", "mp4", "wmv", "mpg", "mov", "mpeg", "rm", "swf", "flv", "ram") {
		return "video"
	} else {
		return format
	}
}

func GetImageInfo(path string) (image2.Rectangle, error) {
	var imageInfo image2.Rectangle
	reader, err := os.Open(path)
	if err != nil {
		return imageInfo, err
	}
	defer reader.Close()
	image, _, err := image2.Decode(reader)
	if err != nil {
		return imageInfo, err
	}
	imageInfo = image.Bounds()
	return imageInfo, nil

}

func GetFileMd5(filename string) (string, error) {
	// 文件全路径名
	pFile, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("打开文件失败，filename=%v, err=%v", filename, err)
		return "", err
	}
	defer pFile.Close()
	md5h := md5.New()
	_, err = io.Copy(md5h, pFile)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func GetFilePath() string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(configs.Conf.App.Resource, "/", folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
	}
	return folderPath
}

func GetTmpFilePath() string {

	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(configs.Conf.App.Resource, "/tmp/", folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
	}
	return folderPath
}

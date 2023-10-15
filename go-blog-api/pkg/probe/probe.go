package probe

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	gwc "github.com/jyotiska/go-webcolors"
	"github.com/nfnt/resize"
	_ "golang.org/x/image/webp"
)

func New(path string) Probe {
	var p Probe
	p.Filepath = path
	return p
}
func ImgDownload(url string, store string) (p Probe, err error) {
	path := strings.Split(url, "/")

	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	} else {
		return p, errors.New("wrong download url")
	}
	p.Filename = name

	out, err := os.Create(store + "/" + name)
	defer out.Close()
	if err != nil {
		return
	}
	p.Filepath = store + "/" + name

	defer p.Destroy()

	resp, err := http.Get(url)

	defer resp.Body.Close()
	if err != nil {
		return
	}
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	p.Type = resp.Header.Get("Content-Type")

	size, err := io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		return
	}
	p.Filesize = size

	imagePoint, err := p.GetSize()
	if err != nil {
		return
	}
	imageWidth := imagePoint.X
	imageHeight := imagePoint.Y

	colorweave, err := p.GetColorweave()
	if err != nil {
		return
	}

	p.Width = strconv.Itoa(imageWidth)
	p.Height = strconv.Itoa(imageHeight)
	p.Colorweave = colorweave
	return
}

//以单位获取图片的大小
func (p *Probe) FilesizeByUnit(unit string) (size float64, err error) {
	sizeUnit := sizeMap[unit]

	fi, err := os.Stat(p.Filepath)
	if err != nil {
		return
	}

	size = float64(fi.Size()) / sizeUnit

	size, err = strconv.ParseFloat(fmt.Sprintf("%.2f", size), 64)
	if err != nil {
		return
	}
	return
}

//图片修改时间
func (p *Probe) FileModifiedTime() (int64, error) {
	f, e := os.Stat(p.Filepath)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

//重命名图片
func (p *Probe) Rename(to string) error {
	return os.Rename(p.Filepath, to)
}

//删除图片
func (p *Probe) Destroy() error {
	return os.Remove(p.Filepath)
}

//判断图片是否存在
func (p *Probe) IsExist() bool {
	_, err := os.Stat(p.Filepath)
	return err == nil || os.IsExist(err)
}

//获取图片长
func (p *Probe) GetHeight() (width int, err error) {
	reader, err := os.Open(p.Filepath)
	if err != nil {
		return
	}
	defer reader.Close()

	image2, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	return image2.Bounds().Size().Y, nil
}

//获取图片的宽度
func (p *Probe) GetWidth() (width int, err error) {
	reader, err := os.Open(p.Filepath)
	if err != nil {
		return
	}
	defer reader.Close()
	image2, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	return image2.Bounds().Size().X, nil
}

//同时获取图片的长宽
func (p *Probe) GetSize() (point image.Point, err error) {
	reader, err := os.Open(p.Filepath)
	if err != nil {
		return
	}
	defer reader.Close()
	image2, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	return image2.Bounds().Size(), nil
}

//获取图片的主题色
func (p *Probe) GetColorweave() (colorMap map[string]float64, err error) {
	reader, err := os.Open(p.Filepath)
	if err != nil {
		return
	}
	defer reader.Close()

	image2, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	// Resize the image to smaller scale for faster computation
	image2 = resize.Resize(100, 0, image2, resize.Bilinear)
	bounds := image2.Bounds()

	ColorCounter := make(map[string]int)
	Limit := 5 // Limiting how many colors to be displayed in output
	TotalPixels := bounds.Max.X * bounds.Max.Y

	for i := 0; i <= bounds.Max.X; i++ {
		for j := 0; j <= bounds.Max.Y; j++ {
			pixel := image2.At(i, j)
			red, green, blue, _ := pixel.RGBA()
			RGBTuple := []int{int(red / 255), int(green / 255), int(blue / 255)}
			ColorName := FindClosestColor(RGBTuple, "css3")
			_, present := ColorCounter[ColorName]
			if present {
				ColorCounter[ColorName] += 1
			} else {
				ColorCounter[ColorName] = 1
			}
		}
	}

	// Sort by the frequency of each color
	keys := make([]int, 0, len(ColorCounter))
	for _, val := range ColorCounter {
		keys = append(keys, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	ReverseColorCounter := ReverseMap(ColorCounter)

	colorMap = make(map[string]float64)
	for _, val := range keys[:Limit] {
		size, err := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(val)/float64(TotalPixels)), 64)
		if err != nil {
			break
		}
		colorMap[ReverseColorCounter[val]] = size
	}

	return
}

func FindClosestColor(RequestedColor []int, mode string) string {
	MinColors := make(map[int]string)
	var ColorMap map[string]string

	// css3 gives the shades while css21 gives the primary or base colors
	if mode == "css3" {
		ColorMap = gwc.CSS3NamesToHex
	} else {
		ColorMap = gwc.HTML4NamesToHex
	}

	for name, hexcode := range ColorMap {
		rgb_triplet := gwc.HexToRGB(hexcode)
		rd := math.Pow(float64(rgb_triplet[0]-RequestedColor[0]), float64(2))
		gd := math.Pow(float64(rgb_triplet[1]-RequestedColor[1]), float64(2))
		bd := math.Pow(float64(rgb_triplet[2]-RequestedColor[2]), float64(2))
		MinColors[int(rd+gd+bd)] = name
	}

	keys := make([]int, 0, len(MinColors))
	for key := range MinColors {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return MinColors[keys[0]]
}

func ReverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

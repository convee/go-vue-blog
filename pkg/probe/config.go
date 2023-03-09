package probe

const (
	B float64 = 1 << (10 * iota)
	KB
	MB
)

var sizeMap = map[string]float64{
	"B":  B,
	"KB": KB,
	"MB": MB,
	"b":  B,
	"kb": KB,
	"mb": MB,
	"Kb": KB,
	"Mb": MB,
}

type Probe struct {
	Filepath   string             //图片路径
	Filename   string             //图片名字
	Filesize   int64              //图片大小
	Type       string             //文件类型
	Width      string             //图片宽度
	Height     string             //图片高度
	Colorweave map[string]float64 //主题色
}

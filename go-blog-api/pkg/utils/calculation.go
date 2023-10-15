package utils

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"reflect"
)

// Round 浮点类型保留小数点后n位精度
func Round(f interface{}, n int) (r float64) {
	pow10N := math.Pow10(n)
	switch f.(type) {
	case float32:
		v := reflect.ValueOf(f).Interface().(float32)
		r = math.Trunc((float64(v)+0.5/pow10N)*pow10N) / pow10N
	case float64:
		v := reflect.ValueOf(f).Interface().(float64)
		r = math.Trunc((v+0.5/pow10N)*pow10N) / pow10N
	}
	return r
}

// Thousands 整数千分位
func Thousands(m interface{}) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", m)
}

// ThousandsAndRound 千分位2位小数
func ThousandsAndRound(m interface{}) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.2f", m)
}

// PercentAndRound2 d1/d2*100并保留两位小数
func PercentAndRound2(up, down float64) float64 {
	if down == 0 {
		return 0
	}
	decimalD1 := decimal.NewFromFloat(up)
	decimalD2 := decimal.NewFromFloat(down)
	decimal100 := decimal.NewFromInt(100)
	float64Result, _ := decimalD1.Div(decimalD2).Mul(decimal100).Round(2).Float64()
	return float64Result
}

// PercentAndRound1 d1*100并保留两位小数
func PercentAndRound1(d1 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimal100 := decimal.NewFromInt(100)
	float64Result, _ := decimalD1.Mul(decimal100).Round(2).Float64()
	return float64Result
}

// Fen2Yuan d1/100并保留两位小数
func Fen2Yuan(d1 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimal100 := decimal.NewFromInt(100)
	float64Result, _ := decimalD1.Div(decimal100).Round(2).Float64()
	return float64Result
}

func Percent(d1 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimal100 := decimal.NewFromInt(100)
	float64Result, _ := decimalD1.Mul(decimal100).Float64()
	return float64Result
}

func Div(a, b interface{}) float64 {
	fa := cast.ToFloat64(a)
	fb := cast.ToFloat64(b)
	if fb == 0 {
		return 0
	}
	if fb < 0 {
		return -1
	}
	return fa / fb
}

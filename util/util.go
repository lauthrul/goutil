package util

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/simplifiedchinese"
	"math"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"sort"
	"strconv"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Convert(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func Str2Decimal(str string) decimal.Decimal {
	d, _ := decimal.NewFromString(str)
	return d
}

func SortMap2Slice(mp map[int]interface{}) []interface{} {
	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var ret []interface{}
	for _, k := range keys {
		ret = append(ret, mp[k])
	}
	return ret
}

func EnumNames(obj interface{}) []string {
	rt := reflect.TypeOf(obj)
	if rt.Kind() != reflect.Struct {
		return nil
	}

	mp := map[int]interface{}{}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		index := f.Tag.Get("index")
		field := f.Tag.Get("name")
		if field == "-" {
			continue
		}
		if field == "" {
			field = f.Name
		}
		idx, _ := strconv.Atoi(index)
		mp[idx] = field
	}

	var fields []string
	for _, v := range SortMap2Slice(mp) {
		fields = append(fields, v.(string))
	}
	return fields
}

func EnumValues(obj interface{}) []string {
	rt := reflect.TypeOf(obj)
	if rt.Kind() != reflect.Struct {
		return nil
	}
	rv := reflect.ValueOf(obj)

	mp := map[int]interface{}{}
	for i := 0; i < rv.NumField(); i++ {
		f := rt.Field(i)
		if f.Tag.Get("name") == "-" {
			continue
		}
		index := f.Tag.Get("index")
		format := f.Tag.Get("format")
		if format == "" {
			format = "%s"
		}
		v := rv.Field(i)
		idx, _ := strconv.Atoi(index)
		mp[idx] = fmt.Sprintf(format, v.Interface())
	}

	var values []string
	for _, v := range SortMap2Slice(mp) {
		values = append(values, v.(string))
	}
	return values
}

func ClearConsole() error {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		return cmd.Run()
	default:
		return fmt.Errorf("platform unsupported")
	}
}

func Clamp(x, a, b int) int {
	return int(math.Min(float64(b), math.Max(float64(x), float64(a))))
}

func ClampFloat(x, a, b float64) float64 {
	return math.Min(b, math.Max(x, a))
}

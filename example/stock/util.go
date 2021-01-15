package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/simplifiedchinese"
	"reflect"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
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

func EnumNames(obj interface{}) []string {
	rt := reflect.TypeOf(obj)
	if rt.Kind() != reflect.Struct {
		return nil
	}

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		field := f.Tag.Get("name")
		if field == "-" {
			continue
		}
		if field == "" {
			field = f.Name
		}
		fields = append(fields, field)
	}

	return fields
}

func EnumValues(obj interface{}) []string {
	rt := reflect.TypeOf(obj)
	if rt.Kind() != reflect.Struct {
		return nil
	}
	rv := reflect.ValueOf(obj)

	var values []string
	for i := 0; i < rv.NumField(); i++ {
		f := rt.Field(i)
		if f.Tag.Get("name") == "-" {
			continue
		}
		format := f.Tag.Get("format")
		if format == "" {
			format = "%s"
		}
		v := rv.Field(i)
		values = append(values, fmt.Sprintf(format, v.Interface()))
	}

	return values
}

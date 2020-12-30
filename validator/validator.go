package validator

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

type FnRule func(data string) (bool, error)

var rules map[string]FnRule

func chekRule(data, rule string) (bool, error) {
	if rule, ok := rules[rule]; ok {
		return rule(data)
	}
	return true, nil
}

// RegisterRule register validator rule
func RegisterRule(key string, rule FnRule) {
	if rules == nil {
		rules = make(map[string]FnRule)
	}
	rules[key] = rule
}

// Bind Parse and validate request parameters from http GET or POST.
// data:        request data to be parse. like: a=1&b=2&c=3.
// param:       struct pointer to save parsing result, supported struct tags:
//				* name -	param name, same as struct field name if not specified
//				* rule -	validate rule, like:url,email,ip etc. must RegisterRule before Bind
//				* empty -	allow empty or not [true|false]
//				* default -	default value if empty
//				* min -		min value numeric value or min length for string
//				* max -		max value numeric value or max length for string
//				* msg -		error message return if validate fail
// for example：
//				type ListParam struct {
//					Page     int `empty:"true" default:"1" min:"1" name:"page"`
//					PageSize int `empty:"true" default:"10" min:"1" name:"page_size"`
//					Status   int `empty:"true" default:"2" min:"0" max:"2" name:"status"`
//				}
// return nil if success, otherwise return err
func Bind(data string, param interface{}) error {
	var (
		err  error
		args = map[string]string{}
	)

	arr := strings.Split(data, "&")
	for _, v := range arr {
		p := strings.Index(v, "=")
		if p < 0 {
			continue
		}
		args[v[0:p]] = v[p+1:]
	}
	if len(args) == 0 {
		return fmt.Errorf("empty data")
	}

	rt := reflect.TypeOf(param)
	rv := reflect.ValueOf(param)

	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("param should be ptr")
	}
	rt = rt.Elem()
	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("param should be ptr")
	}
	rv = rv.Elem()

	for i := 0; i < rt.NumField(); i++ {

		f := rt.Field(i)
		v := rv.Field(i)

		var (
			name  string
			rule  string
			empty bool = false
			def   string
			min   int64 = math.MinInt64
			max   int64 = math.MaxInt64
			msg   string
		)

		// name
		name = f.Tag.Get("name")
		if len(name) == 0 {
			name = f.Name // strings.ToLower(f.Name)
		}

		// rule
		rule = f.Tag.Get("rule")

		// empty
		val := f.Tag.Get("empty")
		if len(val) != 0 {
			s := strings.ToLower(val)
			if s == "true" {
				empty = true
			} else if s == "false" {
				empty = false
			} else {
				return fmt.Errorf("empty filed should be true or false")
			}
		}

		// default
		def = f.Tag.Get("default")

		// min
		val = f.Tag.Get("min")
		if len(val) != 0 {
			min, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid min")
			}
		}

		// max
		val = f.Tag.Get("max")
		if len(val) != 0 {
			max, err = strconv.ParseInt(val, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid max")
			}
		}

		// msg
		msg = f.Tag.Get("msg")
		if len(msg) == 0 {
			msg = fmt.Sprintf("invalid %s", name)
		}

		// check default data
		value, ok := args[name]
		if !ok {
			if empty {
				if len(def) == 0 {
					continue
				} else {
					value = def
				}
			} else {
				return fmt.Errorf("%s should not be empty", name)
			}
		}

		// check rule
		if len(rule) != 0 {
			if ok, err = chekRule(value, rule); !ok {
				return fmt.Errorf("%s: %s", name, err)
			}
		}

		// check value
		var res interface{}
		ok = false
		switch f.Type.Kind() {
		case reflect.Bool:
			res, err = strconv.ParseBool(value)
			ok = true
		case reflect.Int:
			res, err = strconv.Atoi(value)
			ok = res.(int) >= int(min) && res.(int) <= int(max)
		case reflect.Int8:
			res, err = strconv.ParseInt(value, 10, 8)
			ok = res.(int8) >= int8(min) && res.(int8) <= int8(max)
		case reflect.Int16:
			res, err = strconv.ParseInt(value, 10, 16)
			ok = res.(int16) >= int16(min) && res.(int16) <= int16(max)
		case reflect.Int32:
			res, err = strconv.ParseInt(value, 10, 32)
			ok = res.(int32) >= int32(min) && res.(int32) <= int32(max)
		case reflect.Int64:
			res, err = strconv.ParseInt(value, 10, 64)
			ok = res.(int64) >= int64(min) && res.(int64) <= int64(max)
		case reflect.Uint:
			res, err = strconv.ParseUint(value, 10, 64)
			ok = res.(uint) >= uint(min) && res.(uint) <= uint(max)
		case reflect.Uint8:
			res, err = strconv.ParseUint(value, 10, 8)
			ok = res.(uint8) >= uint8(min) && res.(uint8) <= uint8(max)
		case reflect.Uint16:
			res, err = strconv.ParseUint(value, 10, 16)
			ok = res.(uint16) >= uint16(min) && res.(uint16) <= uint16(max)
		case reflect.Uint32:
			res, err = strconv.ParseUint(value, 10, 32)
			ok = res.(uint32) >= uint32(min) && res.(uint32) <= uint32(max)
		case reflect.Uint64:
			res, err = strconv.ParseUint(value, 10, 64)
			ok = res.(uint64) >= uint64(min) && res.(uint64) <= uint64(max)
		case reflect.Float32:
			res, err = strconv.ParseFloat(value, 32)
			ok = res.(float32) >= float32(min) && res.(float32) <= float32(max)
		case reflect.Float64:
			res, err = strconv.ParseFloat(value, 64)
			ok = res.(float64) >= float64(min) && res.(float64) <= float64(max)
		case reflect.String:
			res, err = value, nil
			l := utf8.RuneCountInString(value)
			ok = l >= int(min) && l <= int(max)
		default:
			continue // unsupported type
		}

		if err != nil {
			return fmt.Errorf(msg)
		}

		if ok {
			v.Set(reflect.ValueOf(res))
		} else {
			return fmt.Errorf("%s out of range [%d-%d]", name, min, max)
		}
	}

	return nil
}

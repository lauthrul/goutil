package validator

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBind(t *testing.T) {
	RegisterRule("odd", func(data string) (bool, error) {
		v, err := strconv.Atoi(data)
		if err != nil {
			return false, err
		}
		if v%2 != 0 {
			return true, nil
		}
		return false, fmt.Errorf("not odd")
	})

	RegisterRule("even", func(data string) (bool, error) {
		v, err := strconv.Atoi(data)
		if err != nil {
			return false, err
		}
		if v%2 == 0 {
			return true, nil
		}
		return false, fmt.Errorf("not even")
	})
	param := struct {
		FieldA int    `name:"a" rule:"odd"`
		FieldB int    `name:"b" rule:"even"`
		FieldC int    `name:"c" min:"1" max:"16"`
		FieldD string `name:"d" min:"1" max:"16"`
		FieldE string `empty:"true" default:"I'm default value"`
	}{}

	data := []string{
		"a=1&b=2&c=10&d=string&e=hello, world",               // pass
		"a=2&b=2&c=10&d=string&e=hello, world",               // a fail
		"a=1&b=1&c=10&d=string&e=hello, world",               // b fail
		"a=1&b=2&c=17&d=string&e=hello, world",               // c fail
		"a=1&b=2&c=10&d=string string string&e=hello, world", // d fail
		"a=1&b=2&c=10&d=string&e=hello, world",               // pass
		"a=1&b=2&c=10&d=string",                              // pass
		"a=1&b=2&c=10&d=string&FieldE=hello, world",          // pass
		"a=1&b=2&c=10", // d fail
	}
	for _, d := range data {
		if err := Bind(d, &param); err != nil {
			t.Logf("× [%s] | %s", d, err.Error())
		} else {
			t.Logf("√ [%s] | %+v", d, param)
		}
	}
}

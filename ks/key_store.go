package ks

import (
	"github.com/fatih/camelcase"
	"reflect"
	"strings"
)

func InitKeyName(ks interface{}, upper bool, prefix ...string) {
	v := reflect.ValueOf(ks)
	tt := reflect.TypeOf(ks)
	tt = tt.Elem()
	p := ""
	if len(prefix) > 0 {
		p = prefix[0]
	}
	if p != "" && !strings.Contains(p, "_") {
		p += "_"
	}
	for i := 0; i < tt.NumField(); i++ {
		value := strings.ToLower(strings.Join(camelcase.Split(tt.Field(i).Name), "_"))
		if upper {
			value = strings.ToUpper(value)
		}
		value = strings.Trim(value, "_")
		if p != "" {
			value = p + value
		}
		v.Elem().Field(i).SetString(value)
	}
}

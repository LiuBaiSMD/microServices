package util

import (
	"reflect"
)
func GetType(c interface{}) string{
	return reflect.TypeOf(c).String()
}

func GetTypes(c ...interface{}) []string{
	var res []string
	for _, k := range c{
		t := reflect.TypeOf(k).String()
		res = append(res, t)
	}
	return res
}
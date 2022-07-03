package util

import (
	"reflect"
)

// NilCheck return C if C != nil , else A
func NilCheck(c, a interface{}) interface{} {
	if c == nil ||
		(reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()) ||
		c == "" {
		return a
	}
	return c
}

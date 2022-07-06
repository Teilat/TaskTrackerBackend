package util

import (
	"reflect"
)

// NilCheck return c if c != nil , else a
func NilCheck(c, a any) any {
	if c == nil ||
		(reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()) ||
		c == "" {
		return a
	}
	return c
}

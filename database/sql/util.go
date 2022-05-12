package sql

import (
	"github.com/google/uuid"
	"reflect"
)

func NilCheck(c, a interface{}) interface{} {
	if c == nil ||
		(reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil()) ||
		c == uuid.Nil ||
		c == "" {
		return a
	}
	return c
}

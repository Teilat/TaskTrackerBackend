package sql

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilCheck(t *testing.T) {

	uuid1, _ := uuid.Parse("c35392ea-01f4-44a6-95d6-15b458cce664")
	uuid2, _ := uuid.Parse("1116a36e-4221-49b2-af11-7773055390bb")
	str := "asfd"

	tests := []struct {
		name string
		a    interface{}
		b    interface{}
		want interface{}
	}{
		{name: "string1", a: "", b: "string", want: "string"},
		{name: "string2", a: "", b: "", want: ""},
		{name: "string3", a: "string", b: "betterString", want: "string"},

		{name: "uuid1", a: uuid.Nil,
			b:    uuid1,
			want: uuid1},
		{name: "uuid2", a: uuid.Nil,
			b:    uuid.Nil,
			want: uuid.Nil},
		{name: "uuid3", a: uuid.Nil,
			b:    nil,
			want: nil},
		{name: "uuid4", a: uuid1,
			b:    uuid2,
			want: uuid1},
		{name: "struct1", a: struct {
			a string
			b int
			c *string
		}{"sdf", 35, &str},
			b: nil,
			want: struct {
				a string
				b int
				c *string
			}{"sdf", 35, &str}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NilCheck(tt.a, tt.b))
	}
}

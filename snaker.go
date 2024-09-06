package retag

import (
	"bytes"
	"reflect"
	"unicode"
)

type Snaker string

func (s Snaker) MakeTag(t reflect.Type, fieldIndex int, path string) reflect.StructTag {
	return makeTagWithTransform(string(s), CamelToSnake, t, fieldIndex, path)
}

func CamelToSnake(src string) string {
	// Dumb implementation
	var b bytes.Buffer
	for i, r := range src {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			r = unicode.ToLower(r)
		}
		b.WriteRune(r)
	}
	return b.String()
}

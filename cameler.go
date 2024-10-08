package retag

import (
	"reflect"

	"github.com/webx-top/com"
)

type Cameler string

func (c Cameler) MakeTag(t reflect.Type, fieldIndex int, path string) reflect.StructTag {
	return makeTagWithTransform(string(c), com.CamelCase, t, fieldIndex, path)
}

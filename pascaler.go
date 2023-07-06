package retag

import (
	"reflect"

	"github.com/webx-top/com"
)

type Pascaler string

func (p Pascaler) MakeTag(t reflect.Type, fieldIndex int) reflect.StructTag {
	return makeTagWithTransform(string(p), com.PascalCase, t, fieldIndex)
}

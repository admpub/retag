package retag

import (
	"fmt"
	"reflect"
	"strings"
)

func makeTagWithTransform(tagName string, transform func(string) string, t reflect.Type, fieldIndex int, _ string) reflect.StructTag {
	key := tagName
	field := t.Field(fieldIndex)
	value := field.Tag.Get(key)
	parts := strings.SplitN(value, `,`, 2)
	switch parts[0] {
	case `-`:
		return field.Tag
	case ``:
		parts[0] = transform(field.Name)
	default:
		parts[0] = transform(parts[0])
	}
	value = strings.Join(parts, `,`)
	tag := fmt.Sprintf(`%s:"%s"`, key, value)
	return reflect.StructTag(tag)
}

type Custom struct {
	TagName   string
	Transform func(string) string
}

func (c Custom) MakeTag(t reflect.Type, fieldIndex int, path string) reflect.StructTag {
	return makeTagWithTransform(c.TagName, c.Transform, t, fieldIndex, path)
}

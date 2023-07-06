package retag

import (
	"fmt"
	"reflect"
	"strings"
)

func NewJSONTagValue(field, tagValue string) JSONTagValue {
	return JSONTagValue{FieldName: field, TagValue: tagValue}
}

func NewJSONTagValues(field, tagValue string) *JSONTagValues {
	v := &JSONTagValues{}
	v.Add(field, tagValue)
	return v
}

type JSONTagValue struct {
	FieldName string
	TagValue  string
}

func (f JSONTagValue) MakeTag(t reflect.Type, fieldIndex int) reflect.StructTag {
	field := t.Field(fieldIndex)
	if field.Name == f.FieldName {
		value := field.Tag.Get(`json`)
		parts := strings.SplitN(value, `,`, 2)
		parts[0] = f.TagValue
		tag := fmt.Sprintf(`json:%q`, strings.Join(parts, `,`))
		return reflect.StructTag(tag)
	}
	return field.Tag
}

type JSONTagValues []JSONTagValue

func (f *JSONTagValues) MakeTag(t reflect.Type, fieldIndex int) reflect.StructTag {
	field := t.Field(fieldIndex)
	for _, set := range *f {
		if field.Name == set.FieldName {
			value := field.Tag.Get(`json`)
			parts := strings.SplitN(value, `,`, 2)
			parts[0] = set.TagValue
			tag := fmt.Sprintf(`json:%q`, strings.Join(parts, `,`))
			return reflect.StructTag(tag)
		}
	}
	return field.Tag
}

func (f *JSONTagValues) Add(field, tagValue string) {
	*f = append(*f, JSONTagValue{field, tagValue})
}

func SetJSONTag(v interface{}, field, tagValue string, args ...string) interface{} {
	if len(args) == 0 {
		return Convert(v, NewJSONTagValue(field, tagValue))
	}
	tagValues := NewJSONTagValues(field, tagValue)
	size := len(args)
	var key string
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			key = args[i]
			continue
		}
		tagValues.Add(key, args[i])
	}
	return Convert(v, tagValues)
}

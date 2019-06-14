package goost

import (
	"strings"
	"fmt"
)

type StringArray interface {
	Join(strs AnyString) String
	At(index int) String
	Length() int
	Primitive() []string
	Object
}

type _string_array []String

func (this _string_array)Join(sep AnyString) String {
	s := []string{}
	for _, v := range this {
		s = append(s, v.(*_string).s)
	}
	return NewString(strings.Join(s, getStringValue(sep)))
}

func (this _string_array)At(index int) String {
	return this[index]
}

func (this _string_array)Length() int {
	return len(this)
}

func (this _string_array)Primitive() []string {
	s := []string{}
	for _, v := range this {
		s = append(s, v.(*_string).s)
	}
	return s
}

// interface implementation for Object
func (this _string_array)Equals(obj Object) bool {
	return Object(this) == obj
}

func (this _string_array)ToString() string {
	return fmt.Sprintf("%#v", this)
}
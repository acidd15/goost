package goost

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

// string, String type only
type AnyString interface{}

type String interface {
	RuneAt(index int) rune
	Length() int
	IsEmpty() bool
	Compare(str AnyString) int
	CompareToIgnoreCase(str AnyString) int
	ToLower() String
	ToUpper() String
	TrimSpace() String
	Split(sep AnyString) StringArray
	Primitive() string
	Object
}

type _string struct {
	s string
}

func getStringValue(s AnyString) string {
	switch s.(type) {
	case string:
		return s.(string)
	case String:
		return s.(String).(*_string).s
	default:
		panic("Unsupported type.")
	}
}

func NewString(s AnyString) String {
	v := &_string{}
	v.s = getStringValue(s)
	return v
}

func (this *_string)RuneAt(index int) rune {
	return []rune(this.s)[index]
}

func (this *_string)Length() int {
	return utf8.RuneCountInString(this.s)
}

func (this *_string)IsEmpty() bool {
	return len(this.s) == 0
}

func (this *_string)Compare(str AnyString) int {
	return strings.Compare(this.s, getStringValue(str))
}

func (this *_string)CompareToIgnoreCase(str AnyString) int {
	a := strings.ToLower(this.s)
	b := strings.ToLower(getStringValue(str))
	return strings.Compare(a, b)
}

func (this *_string)ToLower() String {
	return NewString(strings.ToLower(this.s))
}

func (this *_string)ToUpper() String {
	return NewString(strings.ToUpper(this.s))
}

func (this *_string)TrimSpace() String {
	return NewString(strings.TrimSpace(this.s))
}

// TODO performance
func (this *_string)Split(sep AnyString) StringArray {
	s := strings.Split(this.s, getStringValue(sep))

	rs := []String{}

	for _, v := range s {
		rs = append(rs, NewString(v))
	}

	return StringArray(_string_array(rs))
}

func (this *_string)Primitive() string {
	return this.s
}

// interface implementation for Object
func (this *_string)Equals(obj Object) bool {
	return Object(this) == obj
}

func (this *_string)ToString() string {
	return fmt.Sprintf("%#v", this)
}
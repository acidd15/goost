package goost

type Object interface {
	Equals(obj Object) bool
	ToString() string
}

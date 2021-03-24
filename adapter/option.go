package adapter

import "reflect"

type Option struct {
	Order     int
	Name      string
	Optional  bool
	ValueType reflect.Kind
}

type OptionValue struct {
	Option
	Value interface{}
}

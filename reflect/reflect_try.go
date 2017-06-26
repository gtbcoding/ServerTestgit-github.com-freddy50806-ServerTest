package StructBind

import (
	"errors"
	"reflect"
)

func StructBind(arg1 interface{}, arg2 interface{}, dst interface{}) error {
	if reflect.TypeOf(&arg1).Elem().Kind().String() != "struct" {
		return errors.New("arg1 is not a struct")
	}
	if reflect.TypeOf(&arg2).Elem().Kind().String() != "struct" {
		return errors.New("arg2 is not a struct")
	}
	if reflect.TypeOf(&dst).Elem().Kind().String() != "struct" {
		return errors.New("dst is not a struct")
	}
	arg1_Value := reflect.ValueOf
	return nil
}

package StructBind

import (
	"errors"
	"fmt"
	"reflect"
)

func StructBind(arg1 interface{}, arg2 interface{}, dst interface{}) error {
	if reflect.TypeOf(arg1).Kind().String() != "struct" {
		return errors.New("arg1 is not a struct")
	}
	if reflect.TypeOf(arg2).Kind().String() != "struct" {
		return errors.New("arg2 is not a struct")
	}
	if reflect.TypeOf(dst).Kind().String() != "struct" {
		return errors.New("dst is not a struct")
	}
	arg1_Value := reflect.ValueOf(arg1)
	arg1_Type := reflect.TypeOf(arg1)
	arg2_Value := reflect.ValueOf(arg2)
	arg2_Type := reflect.TypeOf(arg2)
	dst_Value := reflect.ValueOf(dst)
	dst_Type := reflect.TypeOf(dst)
	for i := 0; i < dst_Value.NumField(); i++ {
		f := dst_Value.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, dst_Type.Field(i).Name, f.Type(), f.Interface())

	}
	return nil
}

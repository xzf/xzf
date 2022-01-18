package xReflect

import (
	"reflect"
	"errors"
)

/*
use for a struct have multiple similar fields and need to be set
e.g.:
type Obj struct{
	Str1 string
	Str2 string
	Str3 string
	Str4 string
	Str5 string
}
var obj Obj
var sourceValue []string{}
for index,str := range sourceValue{
	xReflect.SetField(xReflect.SetFieldReq{
		Obj:&obj,
		Index:index,
		Value:str,
	})
}
*/

type SetFieldReq struct {
	Obj   interface{}
	Index int
	Value interface{}
}

func SetField(req SetFieldReq) error {
	value := reflect.ValueOf(req.Obj)
	if value.Type().Kind() != reflect.Ptr {
		return errors.New("obj need be ptr")
	}
	valueElem := value.Elem()
	fieldValue := valueElem.Field(req.Index)
	fieldValue.Set(reflect.ValueOf(req.Value))
	return nil
}

func MustSetField(req SetFieldReq) {
	err := SetField(req)
	if err != nil {
		panic(err)
	}
}


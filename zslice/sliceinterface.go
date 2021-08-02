package zslice

import (
	"fmt"
	"reflect"
)

func ToSliceInterface(list interface{}) ([]interface{}, error) {
	reflectValue := reflect.ValueOf(list)
	kind := reflectValue.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("%v cant not convert to []interface{}", kind)
	}
	newInterfaces := make([]interface{}, reflectValue.Len())
	for i := 0; i < reflectValue.Len(); i++ {
		newInterfaces[i] = reflectValue.Index(i).Interface()
	}

	return newInterfaces, nil
}

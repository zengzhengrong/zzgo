package zslice

import (
	"fmt"
	"reflect"
)

func ToSliceInterface(items any) ([]any, error) {
	reflectValue := reflect.ValueOf(items)
	kind := reflectValue.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("%v cant not convert to []interface{}", kind)
	}
	newInterfaces := make([]any, reflectValue.Len())
	for i := 0; i < reflectValue.Len(); i++ {
		newInterfaces[i] = reflectValue.Index(i).Interface()
	}

	return newInterfaces, nil
}

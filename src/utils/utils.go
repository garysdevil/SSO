package utils

import (
	"fmt"
	"reflect"
)

// 判断 obj 是否在数组或切片里面；判断 obj 是否在map的索引里
func UContain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	fmt.Println(targetValue)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, fmt.Errorf("not in " + reflect.TypeOf(target).Kind().String()) //errors.Error("not in array")
}

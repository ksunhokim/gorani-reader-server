package util

import (
	"reflect"
)

func IsDistinctSlice(slice interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		slice_ := reflect.ValueOf(slice)
		seen := make(map[interface{}]bool)

		for i := 0; i < slice_.Len(); i++ {
			if val, ok := seen[slice_.Index(i).Interface()]; ok && val {
				return false
			}
			seen[slice_.Index(i).Interface()] = true
		}
		return true
	}
	return false
}

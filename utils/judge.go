package utils

import (
	"reflect"
)

// 指针是否为空
func IsPointerNil(data any) bool {
	if data == nil {
		return true
	}

	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Ptr && value.IsNil() {
		return true
	}

	return false
}

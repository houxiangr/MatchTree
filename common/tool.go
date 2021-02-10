package common

import (
	"reflect"
	"strconv"
)

func TransferInterfaceToString(source interface{}) string {
	switch reflect.TypeOf(source).Kind() {
	case reflect.String:
		return source.(string)
	case reflect.Int:
		return strconv.Itoa(source.(int))
	default:
		return ""
	}
}

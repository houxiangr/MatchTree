package common

import (
	"reflect"
	"strconv"
)

func TransferInterfaceToString(source interface{}) (string,error) {
	switch reflect.TypeOf(source).Kind() {
	case reflect.String:
		return source.(string),nil
	case reflect.Int:
		return strconv.Itoa(source.(int)),nil
	default:
		return "",NotMatchInterfaceType
	}
}

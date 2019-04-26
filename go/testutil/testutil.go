package testutil

import (
	"reflect"
)

type Case struct {
	Input, Output interface{}
	Message       string
}

func Equal(expect, actually []int) bool {
	return reflect.DeepEqual(expect, actually)
}

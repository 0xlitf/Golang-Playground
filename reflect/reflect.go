package main

import (
	"fmt"
	"reflect"
)

//https://stackoverflow.com/questions/43059653/golang-interfacenil-is-nil-or-not
//https://www.v2ex.com/t/426113
func golang_interfacenil_is_nil_or_not() {
	var a = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	if a == nil {
		fmt.Printf("a is nil\n")
	}else {
		fmt.Printf("a is not nil\n")
	}

	var val interface{} = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(val), reflect.ValueOf(val))
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}
}

func main() {
	golang_interfacenil_is_nil_or_not()
}

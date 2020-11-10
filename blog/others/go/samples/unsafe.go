package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var value int64 = 5
	var ptr1 = &value
	// A pointer of type unsafe.Pointer can override the type system of Go.
	// This is unquestionably fast but it can also be dangerous if used incorrectly
	// or carelessly
	var ptr2 = (*int32) (unsafe.Pointer(ptr1))
	fmt.Println(*ptr2) // 5
	*ptr1 = 5434123412312431212
	fmt.Println(value) // 5434123412312431212
	fmt.Println(*ptr2) // -930866580
}
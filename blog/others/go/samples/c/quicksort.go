package main

// #cgo CFLAGS: -I${SRCDIR}/quicksortLib
// #cgo LDFLAGS: ${SRCDIR}/quicksort.a
// #include<stdio.h>
// #include<stdlib.h>
// #include<quicksort.h>
import "C"
import "fmt"
// import "unsafe"

func main() {
	var arr = []C.int{20, 4, 30, 99,3}
	C.quicksort((*C.int)(&arr[0]),C.int(0),C.int(4))
	fmt.Println(arr)

	var newArr = []C.int{99,88,77,66,55,44,33,22,11}
	C.quicksort((*C.int)(&newArr[0]),C.int(0),C.int(8))
	fmt.Println(newArr)

	var a, b, result int
	a = 10
	b = 20
	result = (int )(C.add(C.int(a), C.int(b)))
	fmt.Println(result)
	
}
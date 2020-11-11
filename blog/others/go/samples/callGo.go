package main

import "fmt"
import "unsafe"

//#cgo CFLAGS: -I:${SRCDIR}/calClib
//#cgo LDFLAGS: ${SRCDIR}/callC.a
//#include <stdlib.h>
//#include<stdio.h>
//#include<callC.h>
import "C"

func main(){

	fmt.Println("Going to call a C function!")
	C.cHello()
	myMessage := C.CString("Message: I am from GO!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All systems are OK!")
}
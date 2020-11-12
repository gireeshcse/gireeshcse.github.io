### Compiler

```
go tool compile program.go
// creates archive file instead of object file
go tool compile -pack program.go // program.a (ar archive file)
```
Generates object file (program.o) which is
* machine code in relocatable format which is not directly executable
* requires low memory during linking phase

```
package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats){
	runtime.ReadMemStats(&mem);
	fmt.Println("mem.Alloc: ",mem.Alloc)
	fmt.Println("mem.TotalAlloc",mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc",mem.HeapAlloc)
	fmt.Println("mem.NumGC",mem.NumGC)
	fmt.Println("-----")
}

func main(){
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte,500000)
		if s == nil {
			fmt.Println("Operation failed")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}
```
To get even more detailed output about the way the Go garbage collector operates is by using following 
```
GODEBUG=gctrace=1 go run garbageCollector.go
```

Algorithm used for garbage collection is **tricolor mark-and-sweep algorithm.**

```
package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i , j int
}

func main(){

	var N = 4000000
	var structure []data
	// map for storing all our pointers as integers
    myMap := make(map[int] * int)
	// map for storing values without pointers
	myMapPlain := make(map[int] int)
	// spliting the map into maps known as sharding
	splitMap := make([]map[int] int, 200)

	for i:= range splitMap {
		splitMap[i] = make(map[int] int)
	}

	for i := 0;i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
		myMap[i] = &value
		myMapPlain[i] = value
		splitMap[i%200][i] = value
	}

	runtime.GC()
	// preventing the garbage collector from garbage collecting the
	// structure variable too early, as it is not referenced or used outside
	// of the for loop.
	fmt.Println(structure[0])
	fmt.Println(*myMap[1])
	fmt.Println(myMapPlain[2])
	fmt.Println(splitMap[0][0])

}
```

maps slow down the Go garbage collector whereas slices collaborate much better with it

#### Unsafe code

Code that bypasses the type security and memory security of code.(Mostly pointers)

```
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
```

Many low-level packages, such as runtime , syscall , and os , constantly use the unsafe package.

### Calling C code from Go program

File:cgo.go 

```
package main

// #include<stdio.h>
// void callC(){
// 	printf("Hello World!From C Programming\n");
// }
import "C"
import "fmt"
func main(){
	fmt.Println("Go statement")
	C.callC()
	fmt.Println("From Go program!")
}
```
To Execute 
```
go run cgo.go
```

File: callC.h

```
#ifndef CALLC_H
#define CALLC_H

void cHello();
void printMessage(char* message);

#endif
```

File: callC.c

```
#include<stdio.h>
#include "callC.h"

void cHello(){
    printf("Hello from C!");
}

void printMessage(char* message){
    printf("Message from Go : %s\n",message);
}
```

File: callGo.go

```
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
```
To Execute
```
gcc -c *.c
ar as callC.a *.o
go run callC.go
```

### Calling go functions from C

Each Go function that will be called by the C code needs to be exported first. This means that you should put a comment line starting with //export before its implementation.

Note: no spaces after //

File: usedbyC.go

```
package main

import "C"
import "fmt"

//export HelloGo
func HelloGo(){
	fmt.Println("Hello from GO!")
}

//export Multiply
func Multiply(a,b int) int{
	return a*b
}

func main(){

}
```

Generate a C shared library from the Go code by executing the following command:

```
go build -o usedbyC.o -buildmode=c-shared usedbyC.go
```


Generates the following files

```
usedbyC.h  
usedbyC.o
```

File: willusego.c

```
#include<stdio.h>
#include "usedbyC.h"

int main(int argc, char **argv){
    GoInt x = 12;
    GoInt y = 14;

    printf("About to call go function !");
    HelloGo();

    GoInt result = Multiply(x,y);

    printf("result : %d\n",(int)result);
    printf("It worked\n");
    return 0;
}
```


```
 gcc -o willusego willusego.c ./usedbyC.o
 ./willusego

 //output
Hello from GO!
About to call go function !result : 168
It worked
```

### defer

The defer keyword postpones the execution of a function until the surrounding function returns, which is widely used in file input and output operations because it saves you from having to remember when to close an opened file

It is very important to remember that deferred functions are executed in last in, first out (LIFO) order after the return of the surrounding function.

```
package main
import "fmt"

func func1(){
	for i := 3; i > 0; i--{
		defer fmt.Print(i) // 1 2 3
	}
	fmt.Println()
}

func func2(){
	for i := 3; i > 0; i--{
		defer func() {
			fmt.Print(i) // 0 0 0
		}()
	}
	fmt.Println()
}

func func3(){
	for i := 3; i > 0; i--{
		defer func(n int) {
			fmt.Print(n) // 1 2 3
		}(i)
	}
	fmt.Println()
}

func main(){
	fmt.Println("Defer program")
	func1()
	func2()
	func3()
}
```

### Panic and Recover

**panic()** is a built-in Go function that terminates the current flow of a Go program and starts panicking. On the other hand, the recover() function, which is also a built-in Go function, allows you to take back control of a goroutine that just panicked using panic().

```
package main

import "fmt"

func a(){
	fmt.Println("Inside func a");
	defer func (){
		if c := recover(); c != nil{
			fmt.Println("a: recovering from panic",c)
		}
	}()
	fmt.Println("calling b")
	b()
	fmt.Println("Exiting from a") // not executed
}

func b(){
	fmt.Println("Inside func b")
	panic("b: I am in Panic")
	fmt.Println("Exiting from b")// not executed
}

func main(){
	a();
	fmt.Println("Program terminated!")
}
```

Note: function b() knows nothing about function a() ; however,function a() contains Go code that handles the panic condition of function b().

### Unix utilities

#### strace tool

* Command line utility tool which allows to trace system calls and signals.
```
strace ls
```
* can print count time, calls, and errors for each system call when used with the -c command-line option

```
strace -c find /usr/ 1> /dev/null
strace -c go run sample.go
```
As the normal program output is printed in standard output and the output of strace(1) is printed in standard error, the previous command discards the output of the command that is examined and shows the output of strace.

#### dtrace tool

DTrace, allows us to see what happens behind the scenes on a system-wide basis without the need to modify or recompile anything. It also allows us to work on production systems and watch running programs or server processes dynamically without introducing a big overhead.


### Go Environment

```
package main

import (
	"fmt"
	"runtime"
	"strings"
	"strconv"
)

func main(){
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())

	version := runtime.Version()
	major := strings.Split(version,".")[0][2]
	m1, _ := strconv.Atoi(string(major))
	fmt.Println("Major version :",m1)
}
```

Output

```
You are using gc on a amd64 machine
Using Go version go1.15.3
Number of CPUs: 4
Number of Goroutines: 1
Major version : 1
```

To Get list of Go environmental variables

```
go env
```

### Go Assembler

Go assembler, which is a Go tool that allows us to see the assembly language used by the Go compiler.

```
go tool compile -S goEnv.go
```

### Node Trees

Everything in a Go program is parsed and analyzed by the modules of the Go compiler according to the grammar of the Go programming language. The final product of this analysis is a tree that is specific to the provided Go code and represents the program in a different way that is suited to the compiler rather than to the developer.

```
go tool compile -W goEnv.go
```

The -W parameter tells the go tool compile command to print the debug parse tree after the type checking.

**More info**

```
go build -x defer.go
```
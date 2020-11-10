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

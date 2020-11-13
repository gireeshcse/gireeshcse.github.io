* int8 , int16 , int32 , int64 ; and uint8 , uint16 , uint32 , and uint64
* float32 and float64
* complex64 and complex128 . The first one uses two float32 : one for the real part and the other for the imaginary part of the complex number, whereas complex128 uses two float64 .

```
package main

import "fmt"

func main(){
	c1 := 20 + 5i
	c2 := complex(20,5)
	c3 := c1 + c2
	var c4 complex64 = complex64(c3 + c2)
	cZero := c3 - c3

	fmt.Printf("C1 = %T\nC2 = %T\nC3 = %T\nC4 = %T\nCZERO = %T",c1,c2,c3,c4,cZero)
}
```

```
for index := 0; index < 100;index++ {

}

for {

}

anArray = [5] int {-2, -1, 0, 1, 2}
for index, value := range anArray {

}
len(anArray) // 5
twoD = [3][3] string {{"00","01","02"},{"10","11","12"},{"20","21","22"}}
threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

for i := 0; i < len(threeD); i++ {
    v := threeD[i]
    for j := 0; j < len(v); j++ {
        m := v[j]
        for k := 0; k < len(m); k++ {
            fmt.Print(m[k], " ")
        }
    }
    fmt.Println()
}

for _, v := range threeD {
    for _, m := range v {
        for _, s := range m {
            fmt.Print(s, " ")
        }
    }
    fmt.Println()
}
```
* Go arrays are not dynamic
* Arrays are passed by value.

* Slices are passed by reference

```
aSliceLiteral := []int{1, 2, 3, 4, 5}
integer := make([]int, 20) 

for i := 0; i < len(integer); i++ {
    fmt.Println(integer[i])
}

aSliceLiteral = nil // empty existing slice
integer = append(integer, 12345)
integer[1:3] // 2nd and 3rd element of slice
```
* Go initializes the elements of every slice created with make .
* The original slice will be kept in memory for as long as the smaller re-slice exists because the original slice is being referenced by the smaller re-slice.
* Slices have two main properties: **capacity** and **length**
* The capacity of a slice is the current room that has been allocated for this particular slice and can be found with the cap() function. As slices are dynamic in size, if a slice runs out of room, Go automatically doubles its current length to make room for more elements.
* if the length and the capacity of a slice have the same values and you try to add another element to the slice, the capacity of the slice will be doubled whereas its length will be increased by one.
* You should be very careful when using the copy() function on slices because the built-in copy(dst, src) copies the minimum number of len(dst) and len(src) elements.


```
s1 := []int{1,2,3,4}
s2 := [] int {5,6,7,8,9,10,11}

copy(s1,s2) // dstn,src
fmt.Println("S1 ",s1) // [5 6 7 8]
fmt.Println("S2 ",s2) // [5 6 7 8 9 10 11]

s1 := make([][]int, 4)

```


```
package main

import (
    "fmt"
    "sort"
)

type aStructure struct {
    person string
    height int
    weight int
}
func main() {
    mySlice := make([]aStructure, 0)
    mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
    mySlice = append(mySlice, aStructure{"Bill", 134, 45})
    mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
    mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
    mySlice = append(mySlice, aStructure{"Athina", 134, 40})
    fmt.Println("0:", mySlice)

    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i].height < mySlice[j].height
    })
    fmt.Println("<:", mySlice)

    sort.Slice(mySlice, func(i, j int) bool {
        return mySlice[i].height > mySlice[j].height
    })
    fmt.Println(">:", mySlice)
}
```

Appending an Array

```
s := []int{1, 2, 3}
a := [3]int{4, 5, 6}
ref := a[:]
t := append(s, ref...) // creating new slice
s = append(s, ref...) // existing slice
s = append(s, s...) // s+s
```
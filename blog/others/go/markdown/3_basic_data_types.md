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
* The capacity of a slice is the current room that has been allocated for this particular slice and can be found with the **cap()** function. As slices are dynamic in size, if a slice runs out of room, Go automatically doubles its current length to make room for more elements.
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

**Go map**

* Go map is a reference to a hash table

```
iMap = make(map[string]int)
iMap["k1"] = 12
iMap["k2"] = 13
anotherMap := map[string]int {
    "k1": 12
    "k2": 13
}
delete(anotherMap, "k1")
for key, value := range iMap {
    fmt.Println(key, value) // order is random.
}
_, ok := iMap["doesItExist"]
if ok {
    fmt.Println("Exists!")
} else {
    fmt.Println("Does NOT exist")
}

aMap := map[string]int{}
aMap["test"] = 1
```

Note: If we try to get the value of a map key that does not exist in the map, we will get zero

Calling the same **delete()** statement multiple times does not make any difference and does not generate any warning messages.


**constants**

```
const HEIGHT = 200
const (
    C1 = "C1C1C1"
    C2 = "C2C2C2"
    C3 = "C3C3C3"
)

const s1 = 123
const s2 float64 = 123

var v1 float32 = s1 * 12 // no problem
var v2 float32 = s2 * 12 // fails because s2 and v2 are of different types

```
A Go **type** is a way of defining a new named type that uses the same underlying type as an existing type. This is mainly used for differentiating between different types that might use the same kind of data.

```
type Digit int
type Power2 int

const (
    Zero Digit = iota 
    One
    Two
    Three
    Four
)

const (
    p2_0 Power2 = 1 << iota
    _
    p2_2
    _
    p2_4
    _
    p2_6
)
```

**Pointers**

```
func getPointer(n *int) {
}

func returnPointer(n int) *int {
}

func getPointer(n *int) {
    *n = *n * *n
}

func returnPointer(n int) *int {
    v := n * n
    return &v
}

```

* Pointers allow you to share data, especially between Go functions.
* Pointers can be extremely useful when we want to differentiate between a zero value and a value that is not set.

**Date and Times**

```
fmt.Println("Epoch time:", time.Now().Unix())
t := time.Now()
fmt.Println(t, t.Format(time.RFC3339))
fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())
time.Sleep(time.Second)
t1 := time.Now()
fmt.Println("Time difference:", t1.Sub(t))
formatT := t.Format("01 January 2006")
fmt.Println(formatT)
loc, _ := time.LoadLocation("Europe/Paris")
londonTime := t.In(loc)
fmt.Println("Paris:", londonTime)
```

time.Nanosecond , time.Microsecond , time.Millisecond , time.Minute , and time.Hour

The Go constants for parsing the time are
15 => for hour
04 => for minute
05 => for second
PM => for PM
pm => for pm
11 => for month
Jan => for parsing the three-letter abbreviation used for describing a month, 
2006 => for parsing the year
02 => for parsing the day of the month. 
Monday => for parsing strings that contain a long weekday string
Mon => for the abbreviated version of the weekday

If you use January instead of Jan , you will get the long name of the month instead of its three-letter abbreviation, which makes perfect sense.

```
d, err := time.Parse("15:04", myTime)
if err == nil {
    fmt.Println("Full:", d)
    fmt.Println("Time:", d.Hour(), d.Minute())
} else {
    fmt.Println(err)
}

start := time.Now()
time.Sleep(time.Second)
duration := time.Since(start)
```
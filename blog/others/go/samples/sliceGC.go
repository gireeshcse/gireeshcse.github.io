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
	fmt.Println(splitMap[199][0],splitMap[199][199])
	for i := 0; i < 200; i++ {
		fmt.Println(i,"\t",len(splitMap[i])) // 20000 * 200
	}
	

}
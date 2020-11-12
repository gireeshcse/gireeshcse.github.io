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
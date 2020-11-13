package main

import "fmt"

func main(){
	c1 := 20 + 5i
	c2 := complex(20,5)
	c3 := c1 + c2
	var c4 complex64 = complex64(c3 + c2)
	cZero := c3 - c3

	fmt.Printf("C1 = %T\nC2 = %T\nC3 = %T\nC4 = %T\nCZERO = %T",c1,c2,c3,c4,cZero)

	s1 := []int{1,2,3,4}
	s2 := [] int {5,6,7,8,9,10,11}
	s3 := [] int {5,6,7,8,9,10,11}

	copy(s3,s1)
	fmt.Println("S3 ",s3)
	fmt.Println("S2 ",s2)
	copy(s1,s2)
	fmt.Println("After Copy\nS1 ",s1)
	fmt.Println("S2 ",s2)
}
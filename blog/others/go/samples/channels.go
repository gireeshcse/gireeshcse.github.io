package main

import "fmt"

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go one(c1)
	go second(c2)

	for {
		select {
			case msg, ok := <- c1 :
				fmt.Println(msg)
				if !ok {
					c1 = nil
				}
			case msg, ok := <- c2 :
				fmt.Println(msg)
				if !ok {
					c2 = nil
				}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}

	fmt.Println("\nExiting")
}

func one(msg chan string){
	for i := 0; i < 5 ; i++ {
		msg <- "Channel One"
	}
	close(msg)
}

func second(msg chan string){
	for i := 0; i < 6; i++ {
		msg <- "Channel Two"
	}
	close(msg)
}

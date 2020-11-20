package main 

import (
	"fmt"
	"time"
	"os"
	"path/filepath"
)

func main(){

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

	fmt.Println("No of arguments passed ",len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(filepath.Base(os.Args[0]))

	
	timesArray := [3] string { "20-11-2020 16:50", "16:50","2020-11-20 17:50"} 

	for _, myTime := range timesArray{
		d, err := time.Parse("15:04", myTime)
		if err == nil {
			fmt.Println("Full:", d)
			fmt.Println("Time:", d.Hour(), d.Minute())
		} else {
			fmt.Println(err)
		}
	}
	
}
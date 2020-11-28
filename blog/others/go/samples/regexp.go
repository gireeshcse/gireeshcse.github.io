package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func reg1(line string){
	r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
	if r1.MatchString(line){
		match := r1.FindStringSubmatch(line)
		fmt.Println(match[1]) // 01/Jan/2021:15:05:05 +0530
		d1, err := time.Parse("02/Jan/2006:15:04:05 -0700",match[1])
		if err == nil {
			newFormat := d1.Format(time.Stamp)
			fmt.Println(strings.Replace(line, match[1], newFormat, 1))
		}
	}
}

func findIP(input string) string {
    partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
    grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
    matchMe := regexp.MustCompile(grammar)
    return matchMe.FindString(input)
}


func main(){
	line1 := "Some good thing will happen on [01/Jan/2021:15:05:05 +0530] Always hope for the best"
	line2 := "hope world will be normal by [Jan-01-22:15:05:05 +0530] Best is yet to come"

	reg1(line1) // 01/Jan/2021:15:05:05 +0530

	r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\]`)
	if r2.MatchString(line2){
		match := r2.FindStringSubmatch(line2)
		fmt.Println(match[1]) // Jan-01-22:15:05:05 +0530
		d1, err := time.Parse("Jan-02-06:15:04:05 -0700",match[1])
		if err == nil {
			newFormat := d1.Format(time.Stamp)
			fmt.Println(strings.Replace(line2, match[1], newFormat, 1))
		}
	}
	line3 := "we got a request from the user whose ip is 192.168.2.15. And browser Edge" 
	ip := findIP(line3)
	fmt.Println(ip)
}
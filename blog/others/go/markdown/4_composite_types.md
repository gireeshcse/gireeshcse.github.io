#### Structures

```
type aStructure struct {
    person string
    height int
    weight int
}
var s1 aStructure
// to access s1.person
p1 := aStructure{"fmt", 12, -2}

p1 := aStructure{weight: 12, height: -2}

func createStruct(p string, h, w int) *aStructure{
    if h > 10{
        h = 7
    }

    return &aStructure{p, h, w}
}

s1 := createStruct("Ram", 7, 75)
fmt.Println((*s1).person)
fmt.Println((p1.person)

pS := new(aStructure) //new returns a pointer.its allocated memory zeroed but not initialized.

sP := new([]aStructure) // creates a slice with new that points to nil
```

#### Tuples

```
func retThree(x int) (int, int, int) {
return 2 * x, x * x, -x
}

n1, n2, n3 := retThree(20)

```

#### File handling

* bufio.ReadString 

    - reads a file until the first occurrence of its parameter
    - returns a **byte slice**

```
f, err := os.Open(filename)
if err != nil {
    fmt.Println("error opening the file %s\n", err)
    return
}
defer f.close()

r := bufio.NewReader(f)
for {
    line, err = r.ReadString('\n')
    if err == io.EOF {
        break
    }else if err != nil {
        fmt.Println("error reading the file %s",err)
    }

    data := strings.Fields(line) // splits a string based on whitespace characters
    
}

```

#### Regular Expressions and pattern matching

```
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
}

func findIP(input string) string {
    partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
    grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
    matchMe := regexp.MustCompile(grammar)
    return matchMe.FindString(input)
}
```

### Strings

* read only byte slice


#### rune

* **rune** is an int32 value.(represents Unicode code point)
* generally a string is considered as the collection of runes.
* a rune literal is a character in single quotes.

```
const r1 = '€'
fmt.Println("(int32) r1:", r1)
fmt.Printf("(HEX) r1: %x\n", r1)
fmt.Printf("(as a String) r1: %s\n", r1)
fmt.Printf("(as a character) r1: %c\n", r1)

aString := []byte("Mihalis")
for x, y := range aString {
	fmt.Println(x, y)
	fmt.Printf("Char: %c\n", aString[x])
}
fmt.Printf("%s\n", aString)
```

#### Unicode

unicode.IsPrint() , can help you to identify the parts of a string that are printable using runes

```
const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
for i := 0; i < len(sL); i++ {
	if unicode.IsPrint(rune(sL[i])) {
		fmt.Printf("%c\n", sL[i])
	} else {
		fmt.Println("Not printable!")
	}
}
```

#### strings

```
package main

import (
	"fmt"
	s "strings"
	"unicode"
)
var f = fmt.Println
func main(){
	upper := s.ToUpper("Hello There!")
	lower := s.ToLower("Hello There!")
}
```

```
s.EqualFold("Mihalis", "MIHAlis") // True
s.HasPrefix("Mihalis", "Mi") // True
s.HasSuffix("Mihalis", "is") // True
s.Index("Mihalis", "ha") // 2
s.TrimSpace(" \tThis is a line. \n")
s.TrimLeft(" \tThis is a\t line. \n", "\n\t")
s.TrimRight(" \tThis is a\t line. \n", "\n\t ")
s.Count("Mihalis", "i") // 2
s.Repeat("ab", 5) // ababababab
s.Compare("Mihalis", "MIHALIS")
s.Fields("This is a string!")
s.Split("abcd efg", "") // [a b c d e f g]
s.Replace("abcd efg", "", "_", -1) // _a_b_c_d_ _e_f_g_
s.Replace("abcd efg", "", "_", 4) // _a_b_c_d efg
lines := []string{"Line 1", "Line 2", "Line 3"}
s.Join(lines, "+++") // Line 1+++Line 2+++Line 3
s.SplitAfter("123++432++", "++") // [123++ 432++ ]

trimFunction := func(c rune) bool {
	return !unicode.IsLetter(c)
}
s.TrimFunc("123 abc ABC \t .", trimFunction) // abc ABC
```

### switch

```
switch asString {
	case "1":
		fmt.Println("One!")
	case "0":
		fmt.Println("Zero!")
	default:
		fmt.Println("Do not care!")
}

switch {
	case number < 0:
		fmt.Println("Less than zero!")
	case number > 0:
		fmt.Println("Bigger than zero!")
	default:
		fmt.Println("Zero!")
}

var negative = regexp.MustCompile(`-`)
var floatingPoint = regexp.MustCompile(`\d?\.\d`)
var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point!")
	case email.MatchString(asString):
		fmt.Println("It is an email!")
		fallthrough
	default:
		fmt.Println("Something else!")
}
```
The **fallthrough** keyword tells Go to execute the branch that follows the current one, which, in this case, is the default branch.

### Some Math

```
import (
	"math"
	"math/big"
)
var precision uint = 0
pi := new(big.Float).SetPrec(precision).SetFloat64(0)

```
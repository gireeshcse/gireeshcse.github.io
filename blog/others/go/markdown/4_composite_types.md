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
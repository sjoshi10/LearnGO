# Composite Types 

* Arrays and structs are aggregate types; their value are concatenations of other values in memory. 

### 4.1 Arrays

* An array is a fixed-length sequence of zero or more elements of a particular type. Because of their fixed length, arrays are rarely used directly in GO. Slices, which can grow and shrink, are much more versatile, but to understand slices we must understand arrays first. 

```
var a [3]int             // array of 3 integer
fmt.Println(a[0])        // print the first element
fmt.Println(a[len(a)-1]) // print the last element, a [2]

// Print the indices and elements
var a [3]int = [3]int{22,11,33}

for _,v := range a {

 fmt.Println(v)
}
// 22
// 11
// 33
```

The size of an array is part of its type, so [3]int and [4]int are different types. 
```
q := [3]int{1, 2, 3}
q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
```

* If an array's element type is comparable then the array type is comparable too, so we may directly compare two arrays of tha ttype using the == operator, which reports whether all corresponding elments are equal. 

```
a := [2]int{1, 2}
b := [...]int{1,2}
c := [2]int(1,3)

fmt.Println(a == b, a == c, b ==c) // true false false
d := [3]int{1,2}
fmt.Println(a == d) // compile error: connot compare [2]int == [3]int
```

### 4.2 Slices
* A slice literal looks like an array literal, a sequnce of values separated by commas and surrounded by braces, but size is not given. This implicitly creates an array variables of the right size and yields a slice that points to it. 
`s := []int{0, 1, 2, 3, 4, 5}`

* Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contains the same elements. We can create our own function to compare: 
```
func equal(x, y []string) bool{
  if len(x) != len(y){
      return false
  }
  for i := range x {
    if x[i] != y[i]{
       return false
    }
  
  }
  return true
}
```

* Append to slices
```
package main 

import "fmt"

func main(){

num := []int{1,2,3,4}

for x := range num{
  fmt.Println(num[x])// 1 2 3 4 
}

num = append(num,5)

for x := range num{

  fmt.Println(num[x])// 1 2 3 4 5
}


}
```

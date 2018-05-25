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

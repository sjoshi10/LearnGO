# Composite Types

Arrays and Structs are aggregate types; their values are concatenations of other values in memory. Arrays and Structs are fixed size. In contrast, slices and maps are dynamic data structures that grow as values are added. 


## 4.1 Arrays

An array is a fixed-length sequence of zero or more elements of a particular type. Because of their fixed length, arrays are rarely used directly in Go. Slices, which can grow and shrink, are much more versatile, but to understand slices we must understand arrays first. 


```
var a [3]int                 // array of 3 integers 
fmt.Println(a[0])            // print the first element
fmt.Println(a[len(a)-1])     // print the last element, a[2]

//Print the indices and element.

for i, v:= range a {
	fmt.Printlf("%d %d", i, v)

}


for _,i= range a {
	fmt.Printlf("%d %d", i, v)

}
```

We can use an array literal to initialize  an array with list of values. 

`var q [3]int = [3]int{1,2,3}`

In an array literal, if an ellipsis "..." appears in place of length, the array length is determined by the numbers of initializers. 
`q := [...]int{1 , 2, 3}`


The size of an array must be a constant expression, that is, an expression whose value can be computed as the program is being compiled. 
```
q := [3]int{1,2,3}
q = [4]int{1,2,3,4}  // compile error: connot assign [4]int to [3]int
```
As we'll see, the literal syntax is similar for arrays, slices, maps, and structs. The specific for above is a list of values in order, but it is also possible to specify a list of index and value pairs, like this:
```
type Currency int

const (
	USD Currency = iota
    EUR
    GBP
    RMB
)

symbol :] [...]string{USD: "a", EUR: "b", GBP: "c", RMB: "d"}

fmt.Println(RMB, symbol[RMB] ) // "3 d"
```
In this form, indices can appear in any order and some may be omitted, as before, unspecified values take on the zero value for the element type. For instance, 
```
  r := [...]int{99: -1} // definies an array r with 100 elements, all zero except for the last, which has value -1.
```

If an array's element type is comparable then the array type is comparable too, so we may directly compare two arrays of that type using the == operator. 

This function zeroes the contents of a [32]byte array:

```
func zero(ptr *[32]byte){
	for i := range ptr {
		ptr[i]=0
}
```
Arrays are seldom used as function parameters or results; instead we use slices. 


## 4.2 Slices

Slices represent variable-length sequences whose elements all have the same type. A slice type is written []T, where the elements have type T; it looks like an array type without a size. 


Arrays and slices are intimately connected. A slice is a lightweight data structure that gives access to a subsequence of the elements of an array, which is known as the slice's underlying array. 

A slice has three components: a pointer, a length, and a capacity. The pointer points to the first element of the array that is reachable throught the clice, which is not necessarily the array's first element. The length is the number of slice elements; it can't exceed the capacity, which is usually the number of elements between teh start of the slice and the end of the underlying array.



Since a slice contains a pointer to an element of an array, passing a slice to a function permits the function to modify the underlying array elements. 
```
s := int{0,1,2,3} // This implicitly creates an array variable of the right size and yields a slice that points to it. 
```
Unlike arrays, clices are not comparable, so we cannot use == to test whether two slices contain the same elements. We must do the comparison ourselves unless it's a byte type. 

## 4.2.1 The append Function

The built-in append function appends items to slices:
```
s := []int{1, 2, 3, 4}

	fmt.Println(s)   // 1, 2, 3, 4 
	s = append(s,5)
	fmt.Println(s)  // 1, 2, 3, 4, 5
```

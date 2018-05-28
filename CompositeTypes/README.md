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

### 4.3 Maps

The hash table is one of the mose ingenious and versatile of all data structures. It is an unordered collection of key/value pairs in which all the keys are distinct, and the value associated with a given key can be retreived, updated, or removed using a constant number of key comparisons on the average, no matter how large the hash table. 


* In Go, a map is a reference to a hash table, and a map type is writtent map[K]V, where K and V are the types of it's keys and values. All of the keys in a given map are of the same type, and all of the values are of the same type, but the keys need not be of the same type as the values. 

* The built-in function make can be used to create a map:
  `ages := make(map[string]int)` // mapping from strings to ints
  
We can also use a map literal to create a new map populated with some initial key/value pairs

```
package main

import "fmt"

func main(){

ages := make(map[string]int)

ages["saurab"]=27
ages["shashank"]=24

fmt.Println(ages["saurab"])    //27
fmt.Println(ages["shashank"])  //24

//delete Key value pair
delete(ages,"saurab")

ages["shristi"]=29
ages["anjela"]=31
ages["sudeep"]=33
ages["saurav"]=33
ages["rayna"]=18


//This for loop will go through maps
for name, age := range ages{

  fmt.Printf("%s is %d years old.\n", name, age) 


}


// You can use this to see if the element is in map
age, ok := ages["dibyesh"]

if !ok{
  fmt.Println(age) // 0 
  fmt.Println(ok)  // false
  fmt.Println("element does not exist.")

}

age, ok = ages["rayna"]

fmt.Println(age) // 18
fmt.Println(ok)  // true



}
```

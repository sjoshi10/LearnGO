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

* As with slices, maps cannot be compared to each other the only legal comparison is with nil. You can compare each value by looping through the map. 

### 4.4 Structs

A struct is a ggregate data type that groups together zero or more named values of arbitrary types as a single entity. 
```
type Employee struct{
   ID        int
   Name      string
   Address   string
   DOB       time.Time
   Position  string
   Salary    int
   ManagerID int
}

var dilbert Employee    // instance of an employee
dilbert.Salary = 10000  // accessed using dot notation

position := &dilbert.Position
*postion = "Senior Devloper" // use pointer to change value. 

var favoriteEmployee *Employee= &dilbert //use pointer to struct
favoritEmployee.Salary += 5000  
```

* A named struct type S can't declare a field of the same type S: an aggregate value cannot contain itself. But S may declare a field of the pointer type \*S, which lets us create recursive data structures like linked lists and trees. 
```
type tree struct{
  value       int
  right, left *tree
}
```

* The zero value for a struct is composed of the zero values of each of its fields. 
* The struct type with no fields is called the empty struct, written struct{}. It has size zero and carries no information but may be useful nonetheless. 


###### 4.4.1 Struct Literals

A value of a struct type can be written using a struct literal that specifies values for its fields. There are two ways to use struct literals: 

```
type Point struct{ X, Y int}

p:= Point{1, 2}        // first way
q:= Point{X: 1, Y: 2}  // second way
```

##### 4.4.2 Comparing Structs
If all fields of a struct are comparable, the struct itself is comparable, so two expressions of that type may be compare using == or !=. 


##### 4.4.3 Struct Embedding and Anonymous Fields
Go's unusual struct embedding mechanism lets us use on name struct type as an anonymous field fo another sturct type, providing a convenient syntactic shourcut so that a smiple dot expression like x.f can stand for a chaing of fields like x.d.e.f. 

```
type Point struct{
  X, Y int
}

type Circle struct {
  Center Point
  Radius int
}

type Wheel Struct{
  Circle Circle
  Spokes int
}


var w Wheel
w.Circle.Center.X =8
w.Circle.Center.Y =8
w.Circle.Radius = 5
w.Spokes = 20

```

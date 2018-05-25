# Basic Data Types

Go's types fall into four categories: basic types, aggregate types, reference types, and interface types. 
* Basic types: number, strings, and booleans
* Aggregate types: arrays and structs
* Reference types: pointers, maps, functions, and channels. They all refere to program variables or state indirectly. 
* Interface Types. 

### 3.1 Integers
* Signed numbers are represented in 2's complement form, in which the high-order bit is reserved for the sign of the number. 
* For instance the range of int8 is -128 to 127, where as the range of uint8 is 0 to 255. 
* If the result of arithmetic operation, whether signed or unsigned, has more bits than can be represented in the result type, it is said to overflow. 
* Go also provides the following bitwise binary operators, the first four of which treat their operands as bit patterns with no concept of arithmetic carry or sign: 
```
&  bitwise AND
|  bitwise OR
^  bitwise XOR
&^ bit clear (AND NOT)
<< left shift
>> right shift
```
### 3.2 Floating-Point Numbers
* Go provides two sizes of floating-point numbers, float32 and float64. 
* A float32 provides approximately six decimal digits of precision, wheras a float64 provides about 15 digits; float64 should be preferred for most purposes because float32 computations accumulate error rapidly unless one is quite careful, and the smallest positive integer that cannont be exactly represented as a float32 is not large. 

### 3.5 Strings

* A string is an immutalbe sequence of bytes. 
* The built-in len function returns the number of bytes in a string, and the index operations s[i] retrieves the i-th byte of string s. 
```
s := "hello, world"
fmt.Println(len(s))        // 12
fmt.Println(s[0], s[7])    // "h w"
fmt.Println(s[0:5])        //  "hello"
fmt.Println(s[:5])         // "hello"
fmt.Println(s[7:])         // "world"
fmt.Println(s[:])          // "hello, world"
```
* Strings may be compared with comparison operators like == and <; the comparison is done byte by byte, so the result is the natural lexicographic ordering. 

* A raw string literal is written using backquotes instead of double quotes. Within raw string literal, no escape sequences are processed. 
* Raw string literals are a convenient way to write regular expressions. 

### 3.5.4 Strings and Byte Slices
* Four standard packages are particularly important for manipulating strings: bytes, strings, strconv, and unicode. 
* strings package provides many functions for searching, replacing, comparing, trimming, splitting, and joining strings. 
* The bytes package has functions for manipulating slices of bytes, of type []byte, which share some properties with strings. 
* strconv package provides functions for converting boolean, integer, and floating-point values to and from their string representations, and functions for quoting and unquoting strings. 
* The unicode package provides functions like IsDigit, IsLetter, IsUpper, and IsLower for classifying runes. 

```
package main

import (
 "strings"
 "fmt"
)
func main(){
 name := "Saurab"
 index := strings.LastIndex(name,"a")
 fmt.Println(index)                    // 4
 fmt.Println(name[:index])             //saur
 index = strings.Index(name, "a")
 fmt.Println(index)                    // 1
 fmt.Println(name[:index])             //s

}
```
string.LastIndex(s, "f")

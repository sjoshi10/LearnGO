# Pointers

pointer = address of a variable.

```
x := 1  
p := &x               // p, of type *int, points to x
fmt.Println(*p)       //"1"
*p =  2               // equivaletn to x = 2 
fmt.Println(x)        // "2"
```




example 1:
```
func incr(p *int) int{
	*p++ // increments what p points to; does not change p
	return *p

}

v :=1
incr(&v)               // side effect: v is now 2
fmt.Println(incr(&v))  // "3" (and v is 3)
```





You don't need to explicityly allocate and free memory, but to write efficient programs you still need to be aware of the lifetime of variable. For example, keeping unnecessary pointers to short-lived objects within long-0lived objects, especialy global variables, will prevent the garbage collector from reclaiming the short-livd objects. 

Example:
```
var global * int 

func f(){
	var x int 
	x = 1
	global = &x
}
```
Here x must be heap-allocated because it is still reachable from the variable global after x has returned, despite being declared as a local variable. 

## 2.4.1 Tuple Assignment 

Another form of assignment, known as tuple assignment, allows serveral variables to be assigned at once. 
```
x, y =  y, x // this is an example of x and y swapping values. 
i, j, k = 2, 3, 5
```


## 2.5 Type Declarations 

A type declration defines a new named type that has the same underlying type as an existing type. The named type provides a way to seperate different and perhaps incompatible uses of the underlying type so that they can't be mixed unintentionally. 

"type name underlying-type"

Example:
```
type Celsius float64
type Fahrenheit float64 
```
even thought they are both float64, they cannont be compared or combined in arithmetic expressions. 

A conversion from one type to another is allowed if both have the same the same underlying type, or if both are unnamed pointer types that point to variables of the same underlying types. 



## 2.6.2 Pakage Initialization

To use a imported package it needs to be initialized first. It’s done by init function. 
init function doesn’t take arguments neither returns any value. In contrast to main, identifier init is not declared so cannot be referenced:



## 3.1 Integers

int8 - means it can only use up to 8 bits. Which means it can used numbers from -128 to 127
uint8 - unsigned integers mIeans you can't use negative numbers. 0 to 255
unsigned numbers tend to be used only when their bitwise operator or peculiar arithmetic operators are required, as when implementing bit sets, parsing binary formats, or for hashing and cryptography. They are typically not used for merly non-negative quantities. 

If the result of an arithmetic operation, whether signed or unsigned, has more bits than can be represented in teh result type, it is said to overflow. The high-order bits that do not fit are silently discarded. If the original number is a signed type, the result could be negative if the lefmost bit is aq, as in the int8 example here:

```
var u uint8 = 255
fmt.Println(u, u+1, u*u) // 255 0 1

var i int8 =  127
fmt.Println(i, i+1, i*i) // 127 -128 1
```

In general, an explicit conversion is required to convert a value from one type to another, and binary operator for arithmetic and logic (except shifts) must have operands of the same type. ALthough this occasionally reults in longer expressions, it also eliminates a whole class of problems and makes programs easier to understand. 

```
var apples int32 =1 
var oranges int16 =2 
var compote  int = apples oranges // compile error 
```
This type mismatch can be fixed in serveral ways, most directly by converting everything to a commont type:

`var compote = int(apples) + int(oranges)`


When printing number using the fmt package, we can control the radix and format with the %d, %o, and %x verbs, as shown in this example. 
```
o := 0666
fmt.Printf("%d %[1]o")
```
Revist this!!!!



## 3.2 Floating-Point Numbers


Revist this!!!


## 3.5 Strings
```
s := "hello, world"
fmt.Println(len(s))  // 12
fmt.Println(s[0], s[7]) // h and w
```

Attemting to access a byte outside this range results in a panic. 



Ths ubstring operation s[i:j] yields a new string consisting of the bytes of the original string starting at index i and continuing up to, but not including, the byte at tindex j. The result container j-i bytes. 
```
fmt.Println(s[0:5]) // hello


fmt.Println(s[:5]) // hello
fmt.Println(s[7:]) // world
fmt.Println(s[:])  // hello, world
```

The + operator makes a new string by concatenating two strings. 


`fmt.Println("goodbye" + s[5:]) // goodbye, world `

Strings maybe be compared with comparision operators like == and <; the comparison is done byte by byte, so the result is the natural lexicographic ordering. 
String values are immutable: the byte sequence contained in a string value can never be changed, though of course we can assign a new value to a string variable. 

Immutability means that it is safe for two copies of a string to share the same underlying memory, making it cheap tocopy strings of any length. Similarly, a string s and a substring like s[7:] my safely share the same dat so the substring operation is also cheap. No new memory is allocated in either case. 


## 3.5.1 String Literals

A string value can be written as a string literal, a squence of bytes enclosed in double quotes: "Hello, world

Withing a double-quoted string literal, escape sequences that begin with a backlash \ can be used to insert arbitrary byte values into the string. 


A raw string literal is written `... `, using backquotes instead fo double quotes. Within a raw string literal, no escape sequences are processed; the contents are taken literally, including backlashes and newlines, so a raw string literal may spread over serveral lines in the program source. 



## 3.5.2 - 3.5.3 

Go back and review



## 3.5.4 Strings and Byte Slices


Four standards packages are particularly important for manipulating strings: bytes, strings, strconv, and unicode. 



The Strings package provides many functions for searching, replacing, comparing, trimming, splitting, and joining strings. 


The bytes package has similar functions for manipulating slics of bytes, of type []byte, which share same properties with strings. Because strings are immutable, buiding up string incrementally can involve a lot of allocation and copying. In such cases, it's more efficient to use the bytes.Buffer type, which we'll show in a moment. 

The strconv package provides functions for converting boolean, integer, and floating-point alues to and from their string representations, and functions for quoating and unquoting strings. 

The unicode package provides functions like IsDigit, IsLetter, IsUpper, and IsLoweer for classifying runes. Each functions takes a singe rune argument and returns a boolean. Conversion functions like ToUpper and ToLower convert a rune into the given case if it is letter. 












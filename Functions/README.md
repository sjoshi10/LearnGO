# Functions 
A function lets us wrap a sequence of statements as a unit that can be called from elsewhere in a program, perhaps multiple times.

### 5.1 Function Declarations
```
func name(parameter-list) (result-list){
    body 
}
```
The result list specifies the typs of the values that the function returns. If the function returns one unnamed result or no results at all, parentheses are optional and usually omitted. 

Here are four ways to declare a function with two parameters and one result, all of type int:
```
func add(x int, y int) int {return x + y}
func sub(x, y int) (z int) {return z - y; return}
func first(x int, _ int) int {return x}
func zero(int, int) int {return 0}
```
* Arguments are passed by value, so the function receives a copy of each argument; modifications to the copy do not affect the caller. However, if the argument contains some kind of reference, like a pointer, slice, map, function, or channel, then the caller may be affected by any modifications the function makes to variables indirectly referred to by the argument. 

### 5.2 Recursion
Functions may be recursive, that is they may call themselves, either directly or indirectly. 

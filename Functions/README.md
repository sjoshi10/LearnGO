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

### 5.3 Multiple Return Values
* A function can return more than one result. We have seen many examples of functions from standard packages that return two values, the desired computational result and an error value or boolean that indicates whether the computation worked. 

```
func example(url string)([]string, error){
    resp, err := http.Get(url)
    if err!= nil{
        return nil, err
    }
    
    if resp.StatusCode != http.StatusOK{
        resp.Body.Close()
        return nil, fmt.ErrorF("getting %s: %s",url, resp.Status) 
    
    }

}
```
* The first error is returned unchanged. 
* The second error is augmented with additional context information by fmt.Errorf. 

### 5.4 Errors
Errors are an important part of a package's API or an application's user interface, and failure is just one of several expected behaviors. 

* A function for which failure is an expected behavior returns an additional result, conventionally the last one. If the failure has only one possible cause, the resutl is a boolean, usually called ok. 

```
value, ok := cache.Lookup(key)
if !ok {
        // cache[key] does not exist. 
}
```

* The built-in type error is an interface type. Error ma be nil or non-nil, that nil implies success and non-nil implies failure, and that a non-nil error has an error message string we can obtain by calling it's Error method or print by calling `fmt.Println(err)` or `fmt.Printf("%v", err)`


* go's approach sets it apart from mamy other languages in which failures are reported using exceptions, non ordinary values. Although Go does have an exception mechanism of sorts, it is used only for reporting truly unexpected errors that indicate a bug, not the routine errors that a robust program should be built to expect. 


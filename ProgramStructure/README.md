## 2.1 Names
* If an entity is declared within a function, it is local to that function. If declared outside of a function, however, it is visible in all files of the package to which it belongs. The case of the first letter of a name determines it's visibility across package boundaries. If the name begins with an upper-case letter, it is exported, which means that it is visible and accessible outside of its own package and may be reffered to by other parts of the program, as with Printf in the fmt package. Package names themselves are always in lowercase. 

## 2.2 Declarations 
* There are four major kinds of declarations: var, const , type, and func. 
* The name of each package-level entity is visible not only throughout the source file that ocntains its declration, but throughout all the files of the package. 
* A function declaration has a name, a list of parameter, an optional list of results, and the function body, which contains the statements that define what the function does. 
```
func addtion(a int, b int) int {
  return (a + b)
}
```
## 2.3 Variables 
`var name type = expression`
Either the type or the = expression part may be omitted, but not both. 

* If the type is ommited, it is determined by the initializer expression. If the expression is omitted, the initial value is the zero valuze for the type. 
* The zero-value mechanism ensures that a variable always holds a well-defined value of its type; in Go there is no such thing as an uninitialized variable. 
* It is possible to declare and optionally initialize a set of variables in a single declration, witha  matching list of expresssions. Omitting th tytpe allows decration of multiple variables of different types: 
```
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```
### 2.3.1 Short Variable Declrations 
* Within a function, an alternate form called a short variable declration may be used to declare and initialize local variables. It takes the form `name := expression`, and the type of name is determined by the type of expression. 
* As with var declrations, multiple variables may be declared and initialized in the same short variable declaration. 

`i, j := 0, 1`
 
* But declarations with multiple initializer expressions should be used only when they help readability.
* One subtle but important point: a short variable declaration does not necessarily declare all the variables on its left-hand side. If some of them were already declared in the same lexical block, then the short variable declaration acts like an assignment to those variables. For example:
##### This works
```
in , err := os.Open(inFile)
out, err := os.Create(outFile) 
```
##### Throws compilation error!
```
f, err := os.Open(infile)
f, err := os.Create(outfile)
```

##### 2.3.2 Pointers
* A variable is a piece of storage containing a value. 
* A pointer values is the address of a variable. A pointer is thus the location at which a value is stored. Not every value has an address, but every variable does. 
* If a variable is declared `var x int`, the expression `&x`(adddress of x) yields a pointer to an integer variable, that is, a value of type *int. 

```
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p =2           // equivalent to x=2
fmt.Println(x)  // "2"
```
* Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil. 
```
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // true, false, false
```

* Function can return address of the variable
```
var p = f()

func f() *int{
  v := 1
  return &v
}
```

* Function makes it possible for the function to update the variable that was indirectly passed. 
```
func incr(p *int) int{
  *p++
  return *p
}

v := 1
incr(&v) // v =2 
incr(&v) // v =3
```
##### 2.3.3 The new function
Another way to create a variable is to use the built-in function new. The expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns its address, which is a value of type *.
```
p := new(int) // p = 0
*p = 2        // p =2
```
* The new function is relatively rarely used because the most common unnamed variables are of struct types, for which struct literal syntax is more flexibe. 

##### 2.3.4 Lifetime of Variables 
* The lifetime of a variable is the interval of time during which it exists as the program executes. The liftime of a package-level variable is the entire execution of the program. By contrast, local variables have dynamic lifetimes: a new instance is created each time the declration  statement is executed, and the variable lives on until it becomes unreachable, at which point its storage may be recycled. 

* Because the lifetime of a variable is determinted only by whether or not it is reachable, a local variable may outlive a single iteration of the enclosing loop. 

```
var global *int

func f() {
   var x int 
   x =1 
   global = &x

}
```
Here, x must be heap-allocated because it is still reachable from the variable global after f has returned, despite being declared as a local variable; we say x escapes from f. 

* You don't need to explicitly allocate and free memory, but to write efficient programs you still need to be aware of the lifetime of variables. 

### Tuple Assignment 
* Another form of assignment, known as tuple assignment, allows several variables to be assigned at once. 

    ##### Swapping Values
    x, y = y, x
    a[i], a[j]= a[j], a[i]
    
* 

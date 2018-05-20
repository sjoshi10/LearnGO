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
##### Throw compilation error!
```
f, err := os.Open(infile)
f, err := os.Create(outfile)
```

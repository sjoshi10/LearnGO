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

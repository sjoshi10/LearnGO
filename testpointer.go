package main

import (
  "fmt"
)


func main(){

var i  = 5;
fmt.Println(&i)

p := &i
fmt.Println(p)
fmt.Println(*p)

}

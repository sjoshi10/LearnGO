package main


import (
	"fmt"
)


func callMe(chars []byte, word []byte, length int,  k int){


	if k==0{

		fmt.Println(word)
		return
	}

	for i :=0; i<length; i++{
		fmt.Printf("%d -- %d\n",i,k)

		newword  := append(word, chars[i])
		callMe(chars, newword,length,k-1)

	}

}


func main(){

a :=[]byte{'a', 'b', 'c'}
var word []byte
callMe(a,word,len(a), 3)

}

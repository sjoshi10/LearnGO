package main

import (
 "encoding/json"
 "log"
 "fmt"
)

type Employee struct{
  Name     string `json:"name"`
  Salary   int    `json:"salary"`
  Age      int `json:"age"`
  Skills   []string `json:"skills,omitempty"`
}
type Names struct{
  Name string 
}


func main(){
saurab := Employee{ Name: "Saurab Joshi", Salary: 100000, Age: 27, Skills: []string{"Java", "Docker", "Linux", "Golang", "Ruby", "Python"} }
tom := Employee{ Name: "Tom Brady", Salary: 0, Age: 50 }
brian := Employee{ Name: "Brian Dawkins", Salary: 10000, Age: 40, Skills: []string{"Football"} }


var employees []Employee

employees = append(employees, saurab)
employees = append(employees, tom)
employees = append(employees, brian)


data, err := json.Marshal(employees)
if err != nil{

   log.Fatalf("Json marshaling failed: %s", err)
}

fmt.Printf("%s \n",data)

data, err = json.MarshalIndent(employees, "", "   ")
/*
[{"name":"Saurab Joshi","salary":100000,"age":27,"skills":["Java","Docker","Linux","Golang","Ruby","Python"]},{"name":"Tom Brady","salary":0,"age":50},{"name":"Brian Dawkins","salary":10000,"age":40,"skills":["Football"]}]
*/

if err != nil{

   log.Fatalf("Json marshaling failed: %s", err)
}

fmt.Printf("%s \n",data)
/*
[
   {
      "name": "Saurab Joshi",
      "salary": 100000,
      "age": 27,
      "skills": [
         "Java",
         "Docker",
         "Linux",
         "Golang",
         "Ruby",
         "Python"
      ]
   },
   {
      "name": "Tom Brady",
      "salary": 0,
      "age": 50
   },
   {
      "name": "Brian Dawkins",
      "salary": 10000,
      "age": 40,
      "skills": [
         "Football"
      ]
   }
] 
*/
 
var names []Names

if err := json.Unmarshal(data, &names); err != nil {
  log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(names) //[{Saurab Joshi} {Tom Brady} {Brian Dawkins}]

}
 

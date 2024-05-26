package main

import (
	"fmt"

)

type Person struct {
  Name string
  Age int
}

func NewPerson(name string , age int) *Person {
  return &Person{
    Name: name,
    Age: age,
  }
}

func (p *Person) changeName(name string){
  p.Name = name
}

func main(){
  myPerson := NewPerson("Dhruv", 25)

  myPerson.changeName("umang")
  fmt.Printf("this is  a person name %s age %d\n" , myPerson.Name, myPerson.Age )
  fmt.Printf("this is  a person %+v\n" , myPerson )
  fmt.Println(myPerson)
}

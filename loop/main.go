package main

import "fmt"
import "slices"

func main(){
  animals := [2]string{}
  animals[0] = "dog"
  animals[1] =  "cat"


  fmt.Println(animals)

  insects := []string{}
  

  insects = append(insects, "worm", "dung beetle")
  
  insects = slices.Delete(insects, 0, 1)

  fmt.Println(insects)

}

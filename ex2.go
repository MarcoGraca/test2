package main

import (
  "fmt"
  "strings")
func main(){
  b:= [3]string{"Gate", "Comet", "Pizza"}
  fmt.Println(b)
  for _,num := range b{
    if strings.Contains(num, "e"){
      fmt.Println("Found one! ",string(num[0]))
    }
  }
}

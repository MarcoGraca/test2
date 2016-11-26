package main

import (
  "fmt")
func main(){
	for i:=12; i>=0; i--{
    switch i{
    case 2,4,10:
      continue
    default:
      fmt.Println("i =",i)
    }
  }
}

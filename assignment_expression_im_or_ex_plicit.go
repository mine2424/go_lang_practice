package main

import "fmt"

func main(){
  // look like var number int = 3
  // := works type inference
  //number := 3
  //fmt.Printf("%T",number)

  // uint64 is based on 0,and bool is based on false
  var number64 uint64
  var isbool bool
  fmt.Println(number64, isbool)
}

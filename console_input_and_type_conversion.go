package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func console(){
  //it looks like comand line argment
  scanner := bufio.NewScanner(os.Stdin)
  // input the number
  fmt.Printf("birthDay is: ")
  scanner.Scan()
  //if error output, write this → ex) input, err :=....
  //if not, write this → _
  input, _ := strconv.ParseInt(scanner.Text(),10,64)
  fmt.Printf("your age is %d.\n", 2020 - input)
}

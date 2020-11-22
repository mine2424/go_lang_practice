package main

import "fmt"

func main(){
  // %T is type reference, %v is variable, %t express bool
  // %b,o,d,x is base 2 8 10 16
  // %e is immutable floating number (scientific notation) for example 13.98462873e+00
  // %f is cut scientific notation
  // %g express all number
  // %s is string expression
  // %q is double quoted string expression
  // for example %9q %-9q adjust width console by putting number
  // %.2f %.f is rounded floting number and .'2' in number is rounded position
  // fmt.Sprintf convert variable to string
  fmt.Printf("Hello %b",24)
}

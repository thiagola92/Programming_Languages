package main

import "fmt"

func main() {

  if 1 == 1 {
    fmt.Printf("True\n")
  }

  if 1 == 2 {
    fmt.Printf("True\n")
  } else {
    fmt.Printf("False\n")
  }

  if 1 == 2 {
    fmt.Printf("True\n")
  } else if 2 == 2 {
    fmt.Printf("True 2\n")
  } else {
    fmt.Printf("False\n")
  }

  if false {
    fmt.Printf("\n")
  } else {
    fmt.Printf("false é False\n")
  }

}

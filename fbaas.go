package main

import (
        "strconv"
        "fmt"
        "strings"
)

func fizzbuzz(n int) []string {
   var output []string

   for i := 1; i<=n; i++ {
     if i%3 == 0 {
        if i%5 == 0 {
          output = append(output, "fizzbuzz")
        } else {
          output = append(output, "fizz")
        }
     } else if i%5 == 0 {
        output = append(output, "buzz")
     } else {
        output = append(output, strconv.Itoa(i))
     }
   }
   return output
}

func main() {
  fb := fizzbuzz(20)

  fmt.Printf(strings.Join(fb, "\n" ) + "\n")
}

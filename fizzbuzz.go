package main

import "fmt"

func main() {
  i := 1
  for i <= 50 {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("FizzBuzz");
    } else if i % 5 == 0 {
      fmt.Println("Buzz");
    } else if i % 3 == 0 {
      fmt.Println("Fizz");
    } else {
      fmt.Println(i);
    }
    i++;
  }
}

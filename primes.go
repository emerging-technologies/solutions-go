package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
)

func Is_prime(n uint) bool {
  if n < 2 {
    return false
  }
  
  for i := uint(2); i <= uint(math.Floor(math.Sqrt(float64(n)))); i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func Prime_no(n uint) uint {
  last_prime := uint(2)
  no_primes := uint(1)
  for i := last_prime + 1; no_primes < n; i++ {
    if Is_prime(i) {
      no_primes++
      last_prime = i
    }
  }
  return last_prime
}


func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: %v", os.Args[0])
  }

  for i := 1; i < len(os.Args); i++ {
    if u, err :=  strconv.ParseUint(os.Args[i], 10, 64); err == nil {
      fmt.Println("Prime", u, "is", Prime_no(uint(u)))
    } else {
      fmt.Println("Error:", os.Args[i], "is not a number");
    }
  }
}
 

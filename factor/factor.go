package main

import "fmt"
import "os"
import "strconv"

func main() {
  var n int
  nargs := len(os.Args)
  if nargs < 2 {
    error(1, "need at least one argument")
  }

  n, _ = strconv.Atoi(os.Args[1])
  if n < 2 {
    error(2, "n must be greater than or equal to 2")
  }

  if IsPrime(n) {
    fmt.Printf("%d is a prime number\n", n)
  } else {
    fmt.Printf("%d = %s\n", n, Factor(n))
  }
}

func error(errno int, msg string) {
  fmt.Fprintf(os.Stderr, "E%04d: %s\n", errno, msg)
  os.Exit(errno)
}

func IsPrime(n int) bool {
  if n < 2 {
    return false
  }
  return n == minimumPrimeFactor(n)
}

func Factor(n int) string {
  var s, sep string
  for ; n > 1; {
    f := minimumPrimeFactor(n)
    n /= f
    s += sep + strconv.Itoa(f)
    sep = " * "
  }
  return s
}

func nthOddNumber(n int) int {
  return 2 * n + 1
}

// n must be greater than or equal to 2
func minimumPrimeFactor(n int) int {
  if n % 2 == 0 {
    return 2
  }

  n1, m1 := n, n
  i := 1

  odd := nthOddNumber(i) // a sequence of 3, 5, 7, 9, 11 ...
  q := n / odd
  r := n - odd * q
  //fmt.Printf("%d = %d * %d + %d\n", n1, odd, q, r)

  for  ;  r != 0 && odd * odd <= n; {
    m1 = m1 - 2 * q
    n1 = m1 + r
    i++
    odd = nthOddNumber(i)
    q = n1 / odd
    r = n1 - odd * q
    //fmt.Printf("%d = %d * %d + %d\n", n1, odd, q, r)
  }

  if n % odd == 0 {
    return odd
  }

  return n
}

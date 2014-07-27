package main

import "testing"


type testpair struct {
  value int
  fib_value int
}

var tests = []testpair {
  { 0, 0 },
  { 1, 1 },
  { 2, 1 },
  { 3, 2 },
  { 4, 3 },
  { 5, 5 },
  { 6, 8 },
  { 7, 13 },
  { 23, 28657 },
  { 24, 46368 },
  { 25, 75025 },
}

func TestFibonacci(t *testing.T) {
  for _, pair := range tests {
    v := Fib(pair.value)
      if v != pair.fib_value {
        t.Error(
            "For", pair.value,
            "expected", pair.fib_value,
            "got", v,
          )
      }
  }
}

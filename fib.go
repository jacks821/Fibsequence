package main


func Fib(key int) int {
  if key == 0 {
    return 0
  } else if key == 1 || key == 2 {
    return 1
  }
  return Fib(key - 1) + Fib(key - 2)

}

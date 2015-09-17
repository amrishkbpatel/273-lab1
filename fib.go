package main
import "fmt"
func fibseries(n uint) uint {
	// check if n == 0 or n == 1
  if n == 0 || n == 1 {
  	return n
  } else {
  	return fibseries(n-1) + fibseries(n-2)
  }
}

func main() {
	 fmt.Println(fibseries(10))
}

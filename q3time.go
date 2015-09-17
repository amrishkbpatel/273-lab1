package main

import(
  "fmt"
  "time"
)

func sleeptime(n int){
  select {
  case <-time.After(time.Duration(n)*time.Second):
    fmt.Println("End of Sleep")
  }
}

func main(){
  fmt.Println("Start")
  sleeptime(5)
  fmt.Println("End")

  var input string
  fmt.Scanln(&input)
}

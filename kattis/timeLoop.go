package main
import "fmt"

func main()  {
  var t int
  fmt.Scan(&t)
  for i := 1; i <= t; i++ {
    fmt.Print(i)
    fmt.Print(" ")
    fmt.Println("Abracadabra")
  }
}

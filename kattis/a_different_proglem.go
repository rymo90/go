package main
import (
  "fmt"
  "io"
)
func abs(x int64) int64{
  if x < 0 {
    return -x
  }else{
    return x
  }
}

func main()  {
  var v1, v2 int64
  for{
    _, err:=fmt.Scanf("%d%d", &v1, &v2)
    if err ==io.EOF{
      break
    }
    fmt.Printf("%d\n", abs(v1-v2))
  }
}

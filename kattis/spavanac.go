package main

import (
  "fmt"
)
func fortyfem(res []int,F int) []int {
  if res[1] - 45 < 0{
    res[0]--
    res[1] = 60 + res[1] - 45
    }else{
      res[1]-= 45
      }
  if (res[0]< 0){
    res[0] = 23
    }
  return res
}
func main()  {
  var H, M int
  fmt.Scan(&H)
  fmt.Scan(&M)
  res := make([]int,0)
  res = append(res, H ,M)
  var F = 45
  alarm:=fortyfem(res,F)
  for i, v := range alarm {
    if i > 0{
      fmt.Print(" ")
    }
    fmt.Print(v)
  }
  fmt.Println()
}

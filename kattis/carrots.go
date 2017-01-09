package main

import ("fmt"
        "os"
        "bufio"

      )

func main()  {
  var t, n int
  fmt.Scan(&t)
  fmt.Scan(&n)
  contest := make([]string, t)
  for i := 0; i < t; i++ {
    reader:= bufio.NewReader(os.Stdin)
    text, _ :=reader.ReadString('\n')
    contest = append(contest, text)
  }

  fmt.Println(n)

}

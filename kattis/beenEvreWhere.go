package main

import "fmt"

func main() {
	var t, x int
	fmt.Scan(&t)
	s := make([]string, t)
	for i := 0; i < len(s); i++ {
		fmt.Scan(&x)
		fmt.Println(x)

	}
}

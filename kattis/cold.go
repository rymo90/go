package main

import "fmt"

func main() {
	counter := 0
	var n int
	var t int
	fmt.Scan(&n)
	//
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t < 0 {
			counter += 1
		} else {

		}
	}
	fmt.Println(counter)

}

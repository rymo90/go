package main

import "fmt"

func main() {
	var T, n int
	var s string
	fmt.Scan(&T)
	citys := make([][]string, T)
	for i := 0; i < T; i++ {
		fmt.Scan(&n)
		citys[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&s)
			citys[i] = append(citys[i], s)
		}
	}
	fmt.Println(citys)

}

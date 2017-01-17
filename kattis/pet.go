package main

import "fmt"

func main() {
	sum := 0
	var x1, x2, x3, x4 int
	inDex := 0
	for i := 1; i <= 5; i++ {
		s1 := 0
		fmt.Scan(&x1)
		fmt.Scan(&x2)
		fmt.Scan(&x3)
		fmt.Scan(&x4)
		s1 += x1 + x2 + x3 + x4
		if s1 > sum {
			sum = s1
			inDex = i
		}
	}
	fmt.Println(inDex, sum)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	var sum float64
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		thelastone := x % 1e1
		firstwo := x / 1e1
		resultat := math.Pow(float64(firstwo), float64(thelastone))
		sum += resultat
	}
	fmt.Println(int(sum))

}

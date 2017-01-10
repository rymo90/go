package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(x, y, n, first int) string {
	if first%(x) == 0 && first%y == 0 {
		return "FizzBuzz"
	} else if first%x == 0 {
		return "Fizz"
	} else if first%y == 0 {
		return "Buzz"
	} else {
		sista := strconv.Itoa(first)
		return sista
	}

}

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)
	first := 1
	for first <= n {
		svar := fizzbuzz(x, y, n, first)
		fmt.Println(svar)
		first++
	}

}

package main

// check if the numer is even or odd
// working with kattis

import (
	"fmt"
	"strconv"
)

func main() {
	var num, x int
	fmt.Scan(&num)
	s := make([]string, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&x)
		if x%2 == 0 {
			d := strconv.Itoa(x)
			//s = append(s, d)
			s = append(s, d+" is even")
		} else {
			d := strconv.Itoa(x)
			s = append(s, d+" is odd")
		}
	}
	for j := 0; j < len(s); j++ {
		fmt.Println(s[j])
	}

}

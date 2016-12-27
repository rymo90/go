package main

import (
	"fmt"
)

func main() {
	exampel := []int{2, 1, 8, 1, 3, 2}
	two := []int{2, 3, 4, 1, 2, 6}
	for i := 0; i < len(exampel); i++ {
		if exampel[i] > two[i] {

			fmt.Println("högre")
		} else if exampel[i] < two[i] {
			fmt.Println("lägre")
		} else {
			fmt.Println("lika")
		}
	}

}

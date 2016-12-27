package main

import "fmt"

func chest(cet, get []int) []int {
	resultat := make([]int, 0)
	counter := 0
	for i := 0; i < 6; i++ {
		if cet[i] == get[i] {
			counter = 0
			resultat = append(resultat, counter)
		} else {
			for j := 0; j <= 10; j++ {
				if cet[i] > get[i] {
					get[i]++
					counter++
					resultat = append(resultat, counter)
					break
				} else if cet[i] < get[i] {
					get[i]--
					counter--
					resultat = append(resultat, counter)
					break
				}

			}
		}

	}
	return resultat

}

func main() {
	var input int
	cet := []int{1, 1, 2, 2, 2, 8}
	get := make([]int, 0)
	for i := 0; i < len(cet); i++ {
		fmt.Scan(&input)
		get = append(get, input)
	}
	fmt.Println(chest(cet, get))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int64
	fmt.Scan(&N)
	bits := strconv.FormatInt(N, 2)
	bits_number := len(bits)
	number, _ := strconv.ParseUint(bits, 2, bits_number)
	r_number := number - number
	for i := 0; i < bits_number; i++ {
		r_number <<= 1
		r_number |= number & 1
		number >>= 1
	}
	fmt.Printf("%s [%d]\n", strconv.FormatUint(r_number, 2), r_number)

}

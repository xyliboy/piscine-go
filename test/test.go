package main

import (
	"fmt"

	"piscine"
)

func main() {
	arguments := []string{"hello, hi, agamhsoy"}
	fmt.Println(piscine.ConcatParams(arguments))
}

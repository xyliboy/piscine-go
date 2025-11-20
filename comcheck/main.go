package main

import (
	"fmt"
	"os"
)

func main() {
	leksi := os.Args[1:]
	for i := range leksi {
		if leksi[i] == "01" || leksi[i] == "galaxy" || leksi[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}

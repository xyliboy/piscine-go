package main

import (
	"fmt"
	"os"
)

func ComCheck() {
	leksi := os.Args[1:]
	for _, j := range leksi {
		if j == "01" || j == "galaxy" || j == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}

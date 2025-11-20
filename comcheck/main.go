package main

import (
	"fmt"
	"os"
)

func ComCheck() {
	leksi := os.Args[1:]
	for _, i := range leksi {
		if i == "01" || i == "galaxy" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}

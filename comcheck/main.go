package comcheck

import (
	"fmt"
	"os"
)

func ComCheck() {
	leksi := os.Args[1:]
	for i := range leksi {
		if leksi[i] == "01" || leksi[i] == "galaxy" || leksi[i] == "galaxy 01" {
			fmt.Print("Alert!!!")
			break
		}
	}
}

package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

// τυπώνει ένα string με PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// τυπώνει το μήνυμα λάθους όπως στο παράδειγμα
func printError(err error) {
	printString("ERROR: ")
	printString(err.Error())
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	// Αν δεν υπάρχουν arguments → διάβασε από stdin και γράψε σε stdout
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError(err)
			os.Exit(1)
		}
		return
	}

	// Αν υπάρχουν αρχεία → προσπάθησε να τα ανοίξεις ένα-ένα
	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printError(err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, file)
		file.Close()
		if err != nil {
			printError(err)
			os.Exit(1)
		}
	}
}

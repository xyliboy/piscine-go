package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[0]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Διαβάζουμε ΟΛΟ το περιεχόμενο του αρχείου και το τυπώνουμε
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(content))
}

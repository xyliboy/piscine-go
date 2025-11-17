package main

import (
	"fmt"
	"os"
)

func parseCount(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func main() {
	args := os.Args[1:]

	// Περιμένουμε τουλάχιστον: -c <count> <file>
	if len(args) < 3 {
		return
	}

	if args[0] != "-c" {
		return
	}

	count := parseCount(args[1])
	if count <= 0 {
		return
	}

	files := args[2:]
	multiple := len(files) > 1
	hadError := false
	firstPrinted := false

	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			fmt.Printf("%v\n", err)
			hadError = true
			continue
		}

		info, err := f.Stat()
		if err != nil {
			fmt.Printf("%v\n", err)
			f.Close()
			hadError = true
			continue
		}

		if multiple {
			if firstPrinted {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
			firstPrinted = true
		}

		size := info.Size()
		start := int64(0)
		if size > int64(count) {
			start = size - int64(count)
		}

		_, err = f.Seek(start, 0)
		if err != nil {
			fmt.Printf("%v\n", err)
			f.Close()
			hadError = true
			continue
		}

		buf := make([]byte, 1024)
		for {
			n, readErr := f.Read(buf)
			if n > 0 {
				_, _ = os.Stdout.Write(buf[:n])
			}
			if readErr != nil {
				break
			}
		}

		f.Close()
	}

	if hadError {
		os.Exit(1)
	}
}

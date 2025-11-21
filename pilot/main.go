package main

import "fmt"

// Προσθέτουμε μόνο αυτό που λείπει:
type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft int
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1

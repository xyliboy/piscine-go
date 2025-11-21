package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float64
	Agee     int
	Aircraft int
}

const AIRCRAFT1 = 1

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Agee = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

package piscine

func PodiumPosition(podium [][]string) [][]string {
	start := 0
	endd := len(podium) - 1

	for start < endd {
		podium[start], podium[endd] = podium[endd], podium[start]
		start++
		endd--
	}

	return podium
}

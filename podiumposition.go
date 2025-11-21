package piscine

func PodiumPosition(podium [][]string) [][]string {
	start := 0
	end := len(podium) - 1

	for start < end {
		podium[start], podium[end] = podium[end], podium[start]
		start++
		end--
	}

	return podium
}

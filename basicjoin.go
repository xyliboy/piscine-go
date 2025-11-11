package piscine

func BasicJoin(args ...string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		result += args[i]
	}
	return result
}

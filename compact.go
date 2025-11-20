package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	katharo := []string{}
	for _, i := range slice {
		if i != "" {
			katharo = append(katharo, i)
		}
	}
	*ptr = katharo
	return len(katharo)
}

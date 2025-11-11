package piscine

func Concat(str1 string, str2 string) string {
	if str1 == "" {
		return str2
	}
	if str2 == "" {
		return str1
	}
	return str1 + str2
}

package piscine

func ReverseMenuIndex(menu []string) []string {
	correctMenu := make([]string, len(menu))
	for i, item := range menu {
		correctMenu[len(menu)-1-i] = item
	}
	return correctMenu
}

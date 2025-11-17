package piscine

func ConcatParams(args []string) string {
	var leksi string
	for v := range args {
		if v < len(args)-1 {
			leksi = leksi + args[v] + "\n"
		} else {
			leksi = leksi + args[v]
		}
	}
	return leksi
}

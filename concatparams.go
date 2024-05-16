package piscine

func ConcatParams(args []string) string {
	var result string
	for i, arg := range args {
		if i == 0 {
			result = arg
		} else {
			result += "\n" + arg
		}
	}
	return result
}

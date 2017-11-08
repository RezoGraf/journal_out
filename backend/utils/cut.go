package utils


func Cut(text string, limit, limit2 int) (string) {
	runes := []rune(text)
	if len(runes) >= limit {
		str1 := string(runes[limit:limit2])
		return str1
	}
	return text
}

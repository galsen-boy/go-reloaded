package functions

func Article(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	for i := 1; i < len(s); i++ {
		if (s[i-1] == "a" || s[i-1] == "A") && (s[i][0] == 'a' || s[i][0] == 'e' || s[i][0] == 'u' || s[i][0] == 'o' || s[i][0] == 'h') {
			if i == 1 {
				s[i-1] = "An"
			} else {
				s[i-1] = "an"
			}
		} else if (s[i-1] == "an" || s[i-1] == "An" || s[i-1] == "AN" || s[i-1] == "aN") &&
			!(s[i][0] == 'a' || s[i][0] == 'e' || s[i][0] == 'u' || s[i][0] == 'o' || s[i][0] == 'h') {
			if i == 1 {
				s[i-1] = "A"
			} else {
				s[i-1] = "a"
			}
		}
	}
	return s
}

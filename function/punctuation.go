package functions

func Punctuation(s string) []string {
	// separation of commas
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';' || s[i] == '\'' &&
			(i != 0 && !('a' <= s[i-1] && s[i-1] <= 'z' || 'A' <= s[i-1] && s[i-1] <= 'Z' || s[i-1] == '\'') ||
				i != len(s)-1 && !('a' <= s[i+1] && s[i+1] <= 'z' || 'A' <= s[i+1] && s[i+1] <= 'Z' || s[i+1] == '\'')) {
			s = s[:i] + " " + string(s[i]) + " " + s[i+1:]
			i++
		}
	}
	// fmt.Println(s)
	// own strings.Fields with \n
	res := []string{}
	str := ""
	for i, v := range s {
		if v != ' ' {
			str += string(v)
		}
		if str != "" && (v == ' ' || i == len(s)-1 || v == '\n') {
			res = append(res, str)
			str = ""
		}
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] == "" || res[i] == " " {
			res = append(res[:i], res[i+1:]...)
		}
	}
	// punctuation
	res2 := []string{}
	comma := true
	for i := 0; i < len(res); i++ {
		if i != 0 && (res[i] == "." || res[i] == "," || res[i] == "!" || res[i] == "?" || res[i] == ":" || res[i] == ";") {
			res2[len(res2)-1] += res[i]
		} else if res[i] == "'" {
			if comma && i != len(res)-1 {
				comma = false
				if i != len(res)-1 {
					res[i+1] = "'" + res[i+1]
				}
			} else if !comma {
				comma = true
				if i != 0 {
					if len(res2) > 0 {
						res2[len(res2)-1] = res2[len(res2)-1] + "'"
					}
				}
			} else {
				if len(res2) > 0 {
					res2[len(res2)-1] += " " + res[i]
				}
			}
		} else {
			res2 = append(res2, res[i])
		}
	}
	// fmt.Println(res2)
	return res2
}

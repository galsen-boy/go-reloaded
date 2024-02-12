package main

import (
	"fmt"
	functions "go-reloaded/function"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	files := os.Args[1:]
	// files := []string"sample.txt", "result.txt"
	// Argument check
	if len(files) != 2 || !(len(files[0]) >= 4 && files[0][len(files[0])-4:] == ".txt") ||
		!(len(files[1]) >= 4 && files[1][len(files[1])-4:] == ".txt") {
		fmt.Println("Incorrect input")
		return
	}
	content, err := os.ReadFile(files[0])
	if err != nil {
		fmt.Println("Incorrect ReadFile!")
		return
	}
	// Punctuation
	cont := functions.Punctuation(string(content))
	cont = functions.Article(cont)
	for i := 1; i < len(cont); i++ {
		if len(cont[i]) >= 5 && cont[i][:5] == "(hex)" {
			cont[i-1] = functions.Hex(cont[i-1])
			if cont[i-1] == "ERROR" {
				fmt.Println("Incorrect hex")
				return
			}
			cont[i-1] += cont[i][5:]
			cont = append(cont[:i], cont[i+1:]...)
			i--
		} else if len(cont[i]) >= 5 && cont[i][:5] == "(bin)" {
			cont[i-1] = functions.Bin(cont[i-1])
			if cont[i-1] == "ERROR" {
				fmt.Println("Incorrect bin")
				return
			}
			cont[i-1] += cont[i][5:]
			cont = append(cont[:i], cont[i+1:]...)
			i--
		} else if len(cont[i]) >= 3 && cont[i][:3] == "(up" {
			if len(cont[i]) >= 4 && cont[i][:4] == "(up)" {
				cont = CorrectArguments("1", cont, i, 3, strings.ToUpper)
				i--
			} else if len(cont[i]) >= 4 && cont[i][:4] == "(up," {
				cont = CorrectArguments(cont[i+1], cont, i, 3, strings.ToUpper)
				i--
			}
		} else if len(cont[i]) >= 4 && cont[i][:4] == "(low" {
			if len(cont[i]) >= 5 && cont[i][:5] == "(low)" {
				cont = CorrectArguments("1", cont, i, 4, strings.ToLower)
				i--
			} else if len(cont[i]) >= 5 && cont[i][:5] == "(low," {
				cont = CorrectArguments(cont[i+1], cont, i, 4, strings.ToLower)
				i--
			}
		} else if len(cont[i]) >= 4 && cont[i][:4] == "(cap" {
			if len(cont[i]) >= 5 && cont[i][:5] == "(cap)" {
				cont = CorrectArguments("1", cont, i, 4, functions.Cap)
				i--
			} else if len(cont[i]) >= 5 && cont[i][:5] == "(cap," {
				cont = CorrectArguments(cont[i+1], cont, i, 4, functions.Cap)
				i--
			}
		}
	}
	res := ""
	for i, v := range cont {
		if i == len(cont)-1 || v[len(v)-1:] == "\n" {
			res += v
		} else {
			res += v + " "
		}
	}
	if res == "" {
		fmt.Println("testel yeuffou nite way, sois logique")
		return
	}
	os.WriteFile(files[1], []byte(res), 0o666)
}

// check arguments for \n and do function
func CorrectArguments(arg string, cont []string, index, countOfSkip int, T func(s string) string) []string { // return changed content
	count, i := 0, 0
	if arg == "1" {
		count = 1
		for j, q := index-1, 0; j >= 0 && q != count; j-- {
			_, err := strconv.Atoi(cont[j])
			if cont[j] == "/n" || cont[j] == "" {
				index++
			} else if err != nil {
				cont[j] = T(cont[j])
				q++
			}
		}
		upComp, _ := regexp.Match("\\(up\\)", []byte(cont[index]))
		if upComp {
			cont[index-1] += cont[index][4:]
		} else {
			cont[index-1] += cont[index][5:]
		}
		cont = append(cont[:index], cont[index+1:]...)
		return cont
	} else {
		for i < len(arg) && arg[i] != ')' {
			count *= 10
			n, err := strconv.Atoi(string(arg[i]))
			if err != nil {
				return []string{}
			}
			count += n
			i++
		}
		if count == 0 && i == 0 {
			return []string{}
		}
		if index < count {
			return []string{}
		}
		for j, q := index-1, 0; j >= 0 && q != count; j-- {
			_, err := strconv.Atoi(cont[j])
			if cont[j] == "/n" || cont[j] == "" {
				index++
			} else if err != nil {
				cont[j] = T(cont[j])
				q++
			}
		}
		cont[index-1] += cont[index+1][i+1:]
		cont = append(cont[:index], cont[index+2:]...)
	}
	return cont
}

package functions

import (
	"strconv"
	"strings"
)

func Hex(cont string) string {
	i, err := strconv.ParseInt(cont, 16, 64)
	if err != nil {
		return "ERROR"
	}
	return strconv.FormatInt(i, 10)
}

func Bin(cont string) string {
	i, err := strconv.ParseInt(cont, 2, 64)
	if err != nil {
		return "ERROR"
	}
	return strconv.FormatInt(i, 10)
}

func Cap(cont string) string {
	return strings.Title(strings.ToLower(cont))
}

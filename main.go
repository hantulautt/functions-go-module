package functions_go_module

import (
	"math/rand"
	"strings"
)

func UcWords(str string) string {
	return strings.ToTitle(str)
}

func GenerateString(total int) string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, total)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

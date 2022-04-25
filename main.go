package functions_go_module

import (
	"math/rand"
	"strings"
	"time"
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

func Str2Time(str string) int64 {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0
	}
	return t.Unix()
}

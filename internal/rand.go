package internal

import (
	"math/rand"
	"time"
)

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterrunes = []rune("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz123456789_ABCDEFGH")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterrunes[rand.Intn(len(letterrunes))]
	}
	return string(b)
}

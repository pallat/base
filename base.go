package base

import (
	"fmt"
	"strings"
)

const (
	maxRune = 2147483647
	maxBit  = 31
)

func Int(s string) (d int) {
	b := []rune(s)
	for i := 0; i < len(b); i++ {
		d += (int((b[i])-48) << uint(i))
	}
	return
}

func Any(number, base int) rune {
	var r rune
	var n int
	for {
		number, n = number/base, number%base

		if n == 0 && number == 0 {
			return r
		}

		r = rune(n) + r
	}
}

func Code(ch rune, base int) string {
	s := strings.Split(fmt.Sprintf("%b", ch), "")

	b := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	return strings.Join(b, "") + strings.Repeat("0", maxBit-len(b))
}

package base

import (
	"fmt"
	"strings"
)

const (
	maxRune = 2147483647
	maxBase = 2147483648
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

func Code(ch rune, lenght int) string {
	s := strings.Split(fmt.Sprintf("%b", ch), "")

	b := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	return strings.Join(b, "") + strings.Repeat("0", lenght-len(b))
}

func RunesBase(s string, base, max int) ([]rune, int) {
	l := len(s)
	splitted := []string{}

	div, mod := l/max, l%max

	for i := 0; i < div; i++ {
		splitted = append(splitted, s[:max])
		s = s[max:]
	}

	if mod > 0 {
		splitted = append(splitted, s)
		div++
	}

	r := []rune{}

	for i := 0; i < div; i++ {
		r = append(r, Any(Int(splitted[i]), base))
	}

	return r, l
}

func Codes(runes []rune, lenght, max int) string {
	s := ""

	mod := lenght % max

	for k, r := range runes {
		if mod != 0 {
			if k == len(runes)-1 {
				s += Code(r, mod)
				return s
			}
		}
		s += Code(r, max)
	}

	return s
}

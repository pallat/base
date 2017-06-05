package base

import (
	"fmt"
	"strings"
)

const (
	maxByte = 255
	maxBase = 256
	maxBit  = 7
)

func Int(s string) (d int) {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		d += (int((b[i])-48) << uint(i))
	}
	return
}

func Any(number, base int) byte {
	var r byte
	var n int
	for {
		number, n = number/base, number%base

		if n == 0 && number == 0 {
			return r
		}

		r = byte(n) + r
	}
}

func Code(ch byte, lenght int) string {
	s := strings.Split(fmt.Sprintf("%b", ch), "")

	b := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	return strings.Join(b, "") + strings.Repeat("0", lenght-len(b))
}

func BytesBase(s string, base, max int) ([]byte, int) {
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

	r := []byte{}

	for i := 0; i < div; i++ {
		r = append(r, Any(Int(splitted[i]), base))
	}

	return r, l
}

func Codes(bytes []byte, lenght, max int) string {
	s := ""

	mod := lenght % max

	for k, r := range bytes {
		if mod != 0 {
			if k == len(bytes)-1 {
				s += Code(r, mod)
				return s
			}
		}
		s += Code(r, max)
	}

	return s
}

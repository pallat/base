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

// Int converts string of binary format to integer
func Int(s string) (d int) {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		d += (int((b[i])-48) << uint(i))
	}
	return
}

// Any convert decimal number to any base you need
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

// BytesBase get string of binary format and convert to
// number as base you want and please let us know maximun bits
// that you base use in binary
func BytesBase(s string, base, max int) ([]byte, int) {
	if base > maxBase {
		fmt.Println("base can not over", maxBase)
		return []byte{}, 0
	}
	if max > maxBit {
		fmt.Println("base can not over", maxBit)
		return []byte{}, 0
	}
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

// Code converts a byte to string in binary format
// please let us know what is the actually lenght of it
func Code(ch byte, lenght int) string {
	s := strings.Split(fmt.Sprintf("%b", ch), "")

	b := []string{}

	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	return strings.Join(b, "") + strings.Repeat("0", lenght-len(b))
}

// Codes convert bytes to string in binary format
// please let us know what is the actually lenght of it
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

package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntFromStringBinary(t *testing.T) {
	i := Int("1111")
	if i != 15 {
		t.Error("15 is expected but got", i)
	}
}

func TestConvertAnyNumberToAnyBase(t *testing.T) {
	r := Any(1023, 1024)
	if r != rune(1023) {
		t.Error(string(rune(1023)), " is expected but got", r)
	}
}

func TestConvertAnyNumberToAnyBaseMaxRune(t *testing.T) {
	r := Any(maxRune, maxBase)
	if r != rune(maxRune) {
		t.Error(string(rune(maxRune-1)), " is expected but got", r)
	}
	fmt.Println(r)
}

func TestFindMaxBitLenghtOfMaxRune(t *testing.T) {
	s := []rune(patial)

	for i := 0; i < maxRune; i++ {
		if Int(string(s[:i])) >= maxRune {
			fmt.Printf("lenght %d is number %d", i, Int(string(s[:i])))
			break
		}
	}
}

func TestConvertBinaryMaxbitLenght(t *testing.T) {
	example := "1011000110000011110011111101010"

	ch := Any(Int(example), maxBase)

	back := Code(ch, maxBit)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryMaxbitLenghtOverFlow(t *testing.T) {
	example := "1111111111111111111111111111111"

	ch := Any(Int(example), maxBase)

	back := Code(ch, maxBit)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryMaxbitLenghtNotFullyLenght(t *testing.T) {
	example := "10110001100000111100"

	ch := Any(Int(example), maxBase)

	back := Code(ch, 20)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryStringToEncodeRunes(t *testing.T) {
	example := "10110001100000111100111111010101011000110000011110011111101010"

	runes, _ := RunesBase(example, maxBase, maxBit)

	expected := []rune{
		Any(Int("1011000110000011110011111101010"), maxRune),
		Any(Int("1011000110000011110011111101010"), maxRune),
	}

	if !reflect.DeepEqual(expected, runes) {
		t.Errorf("%v is expected but got %v\n", expected, runes)
	}
}

func TestConvertBinaryStringToEncodeRunesAndRevertBack(t *testing.T) {
	example := "10110001100000111100111111010101011000110000011110011111101010"

	runes, lenght := RunesBase(example, maxBase, maxBit)

	s := Codes(runes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

func TestConvertBinaryStringToEncodeRunesAndRevertBackNotFullBlock(t *testing.T) {
	example := "1011000110000011110011111101010101100011000001111001111110101011100"

	runes, lenght := RunesBase(example, maxBase, maxBit)

	s := Codes(runes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

func TestConvertBinaryStringToEncodeRunesAndRevertBackNotFullBlockOverflow(t *testing.T) {
	example := "1111111111111111111111111111111"

	runes, lenght := RunesBase(example, maxBase, maxBit)

	s := Codes(runes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

// var suite = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
var patial = "111111111111111111111111111111111111111111111111111111111111100000000001111111111111111111111111100000000001111111100000000000000000000111111111111111111111111111111111111111111111111111000000000000000000000000000000000000000011111111111111111111111111111111111111111111111000000000000000000001111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"

func TestConvertBinaryString500BitsToEncodeRunesAndRevertBackNotFullBlock(t *testing.T) {
	example := patial

	runes, lenght := RunesBase(example, maxBase, maxBit)

	s := Codes(runes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

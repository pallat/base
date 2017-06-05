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
	r := Any(255, 256)
	if r != byte(255) {
		t.Error(string(byte(255)), " is expected but got", r)
	}
}

func TestConvertAnyNumberToAnyBaseMaxByte(t *testing.T) {
	r := Any(maxByte, maxBase)
	if r != byte(maxByte) {
		t.Error(string(byte(maxByte-1)), " is expected but got", r)
	}
	fmt.Println(r)
}

func TestFindMaxBitLenghtOfMaxByte(t *testing.T) {
	s := []rune(patial)

	for i := 0; i < maxByte; i++ {
		if Int(string(s[:i])) >= maxByte {
			fmt.Printf("lenght %d is number %d", i, Int(string(s[:i])))
			break
		}
	}
}

func TestConvertBinaryMaxbitLenght(t *testing.T) {
	example := "1011000"

	ch := Any(Int(example), maxBase)

	back := Code(ch, maxBit)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryMaxbitLenghtOverFlow(t *testing.T) {
	example := "1111111"

	ch := Any(Int(example), maxBase)

	back := Code(ch, maxBit)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryMaxbitLenghtNotFullyLenght(t *testing.T) {
	example := "101100"

	ch := Any(Int(example), maxBase)

	back := Code(ch, 6)

	if back != example {
		t.Errorf("%v is origin but revert to %v", example, back)
		t.Errorf("%v is origin but revert to %v", len(example), len(back))
	}
}

func TestConvertBinaryStringToEncodeBytes(t *testing.T) {
	example := "10110001100000"

	bytes, _ := BytesBase(example, maxBase, maxBit)

	expected := []byte{
		Any(Int("1011000"), maxByte),
		Any(Int("1100000"), maxByte),
	}

	if !reflect.DeepEqual(expected, bytes) {
		t.Errorf("%v is expected but got %v\n", expected, bytes)
	}
}

func TestConvertBinaryStringToEncodeBytesAndRevertBack(t *testing.T) {
	example := "10110001100000"

	bytes, lenght := BytesBase(example, maxBase, maxBit)

	s := Codes(bytes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

func TestConvertBinaryStringToEncodeBytesAndRevertBackNotFullBlock(t *testing.T) {
	example := "1011000110000011"

	bytes, lenght := BytesBase(example, maxBase, maxBit)

	s := Codes(bytes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

func TestConvertBinaryStringToEncodeBytesAndRevertBackNotFullBlockOverflow(t *testing.T) {
	example := "1111111"

	bytes, lenght := BytesBase(example, maxBase, maxBit)

	s := Codes(bytes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

// var suite = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
var patial = "111111111111111111111111111111111111111111111111111111111111100000000001111111111111111111111111100000000001111111100000000000000000000111111111111111111111111111111111111111111111111111000000000000000000000000000000000000000011111111111111111111111111111111111111111111111000000000000000000001111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"

func TestConvertBinaryString500BitsToEncodeBytesAndRevertBackNotFullBlock(t *testing.T) {
	example := patial

	bytes, lenght := BytesBase(example, maxBase, maxBit)

	s := Codes(bytes, lenght, maxBit)

	if s != example {
		t.Errorf("%v is origin but revert to %v", example, s)
	}
}

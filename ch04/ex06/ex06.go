package main

import (
	"fmt"
	"unicode"
)

func CompressSpace(b []byte) []byte {
	r := []rune(string(b))
	newR := []rune{}
	currentSpace := false
	for i := 0; i < len(r); i++ {
		if currentSpace && unicode.IsSpace(r[i]) {
			continue
		}
		if unicode.IsSpace(r[i]) {
			newR = append(newR, rune(' '))
			currentSpace = true
		} else {
			currentSpace = false
			newR = append(newR, r[i])
		}

	}
	return []byte(string(newR))
}

// IsBlank function is Copied from @orisano
// Check if the first character in the string is blank or not
func IsBlank(b []byte) (isBlank bool, skipMoreCharacters int) {
	switch b[0] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xa0:
		return true, 0
	case 0xe1:
		if len(b) > 2 && b[1] == 0x9a && b[2] == 0x80 {
			return true, 2
		}
	case 0xe2:
		if len(b) <= 2 {
			return false, 0
		}
		x := b[2]

		switch b[1] {
		case 0x80:
			switch {
			case 0x80 <= x && x <= 0x8a:
				return true, 2
			case 0xa8 <= x && x <= 0xa9:
				return true, 2
			case x == 0xaf:
				return true, 2
			}
		case 0x81:
			if x == 0x9f {
				return true, 2
			}
		}
	case 0xe3:
		if len(b) <= 2 {
			return false, 0
		}
		if b[1] == 0x80 && b[2] == 0x80 {
			return true, 2
		}
	}
	return false, 0
}

// CompressSpaceInsideArray returns the result after manipulating inside the array
func CompressSpaceInsideArray(b []byte) []byte {
	i := 0
	k := 0
	for i < len(b) {
		isBlank, skipMore := IsBlank(b[i:])
		if isBlank {
			i = i + skipMore
			if k == 0 || (k > 0 && b[k-1] != ' ') {
				b[k] = ' '
				k = k + 1
			}
		} else {
			b[k] = b[i]
			k = k + 1
		}
		i = i + 1
	}
	return b[:k]
}

func main() {
	s := "Hello, 世界"
	b := []byte(s)
	r := []rune(string(b))
	fmt.Println(r)
	fmt.Println(b)

	s1 := "Hello,      　　　　 世界"
	b1 := CompressSpace([]byte(s1))
	r1 := []rune(string(b))
	fmt.Println(r1)
	fmt.Println(b1)

}

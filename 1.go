package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func hex2Uint(b byte) byte {
	if '0' <= b && b <= '9' {
		return b - '0'
	} else if 'a' <= b && b <= 'f' {
		return b - 'a' + 10
	} else if 'A' <= b && b <= 'F' {
		return b - 'A' + 10
	} else {
		report := fmt.Sprintf("hex2Uint input out of range: %d", b)
		panic(report)
	}
}

func Base64Lookup(n byte) byte {
	if n > 63 {
		report := fmt.Sprintf("Bas64Lookup input out of range: %d", n)
		panic(report)
	} else if n < 26 {
		return byte(n + 'A')
	} else if n < 52 {
		return byte(n - 26 + 'a')
	} else if n < 62 {
		return byte(n - 52 + '0')
	} else if n == 62 {
		return byte('+')
	} else {
		return byte('/')
	}
}

// Challenge 1.1
func Hex2B64(hex string) string {
	s := hex

	// Left-pad the string so that its length is divisible by 3.
	if len(hex)%3 == 1 {
		s = "00" + s
	} else if len(hex)%3 == 2 {
		s = "0" + s
	}

	// The output b64 string has length = 2/3rd of input string.
	result := make([]byte, len(s)*2/3)

	// Convert chunks of 3 hex bytes to chunks of 2 b64 digits.
	for i := len(s) - 1; i >= 0; i -= 3 {
		a := hex2Uint(s[i-2])
		b := hex2Uint(s[i-1])
		c := hex2Uint(s[i])

		d := (a << 2) | ((b & 12) >> 2)
		e := ((b & 3) << 4) | c

		idx := ((i + 1) * 2 / 3) - 1

		result[idx] = Base64Lookup(e)
		result[idx-1] = Base64Lookup(d)
	}

	return string(result)
}

// Challenge 1.2
/*
 * byte2HexChar converts a byte in the range 0-15 to a hex digit
 */
func byte2HexChar(n byte) byte {
	if n > 15 {
		report := fmt.Sprintf("byte2HexChar: uint out of range: %d", n)
		panic(report)
	}
	if n < 10 {
		return byte(n + '0')
	} else {
		return byte(n - 10 + 'a')
	}
}

func XORBuffers(hex1, hex2 string) string {
	if len(hex1) > len(hex2) {
		hex1, hex2 = hex2, hex1
	}

	hex1 = strings.Repeat("0", len(hex2)-len(hex1)) + hex1
	result := make([]byte, len(hex2))

	for i := len(hex2) - 1; i >= 0; i -= 1 {
		u1 := hex2Uint(hex1[i])
		u2 := hex2Uint(hex2[i])
		result[i] = byte2HexChar(u1 ^ u2)
	}

	return string(result)
}

// Challenge 1.3
/*
 * byte2Hex converts an int in the range 0-255 to a pair of hex digits
 * both of which are stored as single bytes.
 */
func byte2Hex(n byte) (result [2]byte) {
	result[0] = byte2HexChar(n & 0x0f)
	result[1] = byte2HexChar((n & 0xf0) >> 4)
	return
}

func repeat(pattern []byte, desiredLen int) string {
	if desiredLen%len(pattern) != 0 {
		panic(
			"repeat: length of byte sequence is not a factor of desired length")
	}

	result := make([]byte, desiredLen)
	for i := 0; i < desiredLen/len(pattern); i++ {
		for j, _ := range pattern {
			result[i*len(pattern)+len(pattern)-1-j] = pattern[j]
		}
	}

	return string(result)
}

func SingleByteXORCipher(h string) {
	// assumption: the string has even length

	for i := 0; i < 256; i++ {
		test := byte2Hex(byte(i))
		cipher := repeat(test[:], len(h))
		// fmt.Println(h, cipher)
		xorOut := XORBuffers(cipher, h)
		s, _ := hex.DecodeString(xorOut)
		fmt.Println(string(s))
	}
}

func main() {
	SingleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
}

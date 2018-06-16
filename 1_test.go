package main

import (
	"testing"
)

func Test1_1(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expect := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	out := Hex2B64(hex)
	if out != expect {
		t.Errorf("1.1 fail\nexpected %s\ngot      %s", expect, out)
	}
}

func Test1_2(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	expect := "746865206b696420646f6e277420706c6179"

	out := XORBuffers(hex1, hex2)
	if out != expect {
		t.Errorf("1.2 fail\nexpected %s\ngot      %s", expect, out)
	}
}

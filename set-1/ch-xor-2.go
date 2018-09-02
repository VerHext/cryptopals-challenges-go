package main

import (
	"fmt"
	"encoding/hex"
)

/*
	The Fixed XOR challenge.
	# https://cryptopals.com/sets/1/challenges/2
	> Write a function that takes two equal-length buffers and produces their XOR combination.

	a = 1c0111001f010100061a024b53535009181c
	b = 686974207468652062756c6c277320657965
 */




func main()  {
	a := []byte("1c0111001f010100061a024b53535009181c");
	b := []byte("686974207468652062756c6c277320657965");
	fmt.Println(xor(a, b))
}



func xor( ar []byte , br []byte ) string {

	if (len(ar) != len(br)) {
		fmt.Println("xor :: buffer len not equals")
		return "";
	}
	a := make([]byte, hex.DecodedLen(len(ar)))
	hex.Decode(a, ar)
	b := make([]byte, hex.DecodedLen(len(br)))
	hex.Decode(b, br)
	dst := make([]byte, len(a));

	for i := 0; i < len(a); i++ {
		dst[i] = (a[i]^b[i]);
	}
	return hex.EncodeToString(dst);
}

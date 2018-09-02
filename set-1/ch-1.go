package main

import (
	b64 "encoding/base64"
	"fmt"
	"encoding/hex"
)

/*
	The base64 challenge.
	# https://cryptopals.com/sets/1/challenges/1
	> Convert hex to base64
	string :: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

 */


func main()  {

	string := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(hexTobase64(string));
	//Result = SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
}


func hexTobase64(input []byte) string{

	dst := make([]byte, hex.DecodedLen(len(input)))
	hex.Decode(dst, input)

	dst2 := b64.StdEncoding.EncodeToString(dst)
	return dst2;
}
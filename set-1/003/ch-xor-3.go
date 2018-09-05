package main

import (
	"fmt"
	"encoding/hex"
)

/*
	The Single-byte XOR cipher challenge.
	# https://cryptopals.com/sets/1/challenges/3
	> The hex encoded string:
				1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
	... has been XOR'd against a single character. Find the key, decrypt the message.
 */


var idealFreqs = []float32{.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241, .0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007}


func main()  {

	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	hexString, _ := hex.DecodeString(string(input))
	xor_crack_key_single(hexString)
}

	func scoreText(a []byte) float32 {
		cts := make([]int, 26)
		for _, ch := range a {
			if 'A' <= ch && ch <= 'Z' {
				ch -= 32
			}
			if 'a' <= ch && ch <= 'z' {
				cts[int(ch)-'a']++
			}
		}
		amount := float32(len(a))
		var score float32
		freqs := make([]float32, 26)
		for i, num := range cts {
			freqs[i] = float32(num) / amount
			score += freqs[i]
		}
		return score
	}

	func  xor_crack_key_single(input []byte)  {
		var maxScore float32
		var maxKey byte
		var maxDecoded []byte
		for i := 0; i <= 255; i++ {
			decoded := xor_encrypt_single_key(input, byte(i))
			score := scoreText(decoded)
			if score > maxScore {
				maxScore = score
				maxKey = byte(i)
				maxDecoded = decoded
			}
		}
		fmt.Println("Result [",string(maxKey),"] :: " ,  string(maxDecoded))
	}


	func xor_encrypt_single_key(a []byte, k byte) []byte {
		res := make([]byte, len(a))
		for i := range a {
			res[i] = a[i] ^ k
		}
		return res
	}



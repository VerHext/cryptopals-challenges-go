package main

import (
	"fmt"
	"encoding/hex"
	"os"
	"bufio"
)

/*
	The single-character XOR challenge.
	# https://cryptopals.com/sets/1/challenges/4
	> One of the 60-character strings in this file has been encrypted by single-character XOR. Find it.

 */

var idealFreqs = []float32{.0817, .0149, .0278, .0425, .1270, .0223, .0202, .0609, .0697, .0015, .0077, .0402, .0241, .0675, .0751, .0193, .0009, .0599, .0633, .0906, .0276, .0098, .0236, .0015, .0197, .0007}

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

func main()  {

	hexLines, _ := readLines("./4.txt")

	for i := 0; i < 60; i++ {
	hexString := hexLines[i]
	hexResult, _ :=	hex.DecodeString(string(hexString))

	xor_crack_key_single(hexResult);
	}
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


func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}



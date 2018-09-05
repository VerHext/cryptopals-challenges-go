package main

import (
	"fmt"
	"encoding/hex"
	"os"
	"bufio"
	"log"
)

/*
	The repeating-key XOR challenge.
	# https://cryptopals.com/sets/1/challenges/5
	> Here is the opening stanza of an important work of the English language:
		In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd
		against I, the next C, the next E, then I again for the 4th byte, and so on.

		(Encrypt a bunch of stuff using your repeating-key XOR function. Encrypt your mail. Encrypt your password file.
		Your .sig file. Get a feel for it. I promise, we aren't wasting your time with this.)


 */



func main()  {

	text := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE");
	fmt.Println(hex.EncodeToString(xor_encrypt(text, key)))

	fmt.Println("---------- [E-Mail] ----------")
	fmt.Println(hex.EncodeToString(xor_encrypt([]byte("support@allesverhext.de"), []byte("secertPassword"))))

	fmt.Println("---------- [Password File] ----------")
	lines, err := readLines("password.txt")
	if err != nil {
		log.Println("readLines: %s", err)
	}
	for _, line := range lines {

		encrypted := xor_encrypt([]byte(line), []byte("youCantHackMe"));
		fmt.Println(hex.EncodeToString(encrypted))
		fmt.Println(string(xor_decrypt(encrypted,[]byte("youCantHackMe"))));
	}


}

func xor_encrypt(text []byte, key []byte) (dest []byte) {
	dest =  make([]byte, len(text))
	keyCount := 0;
	for i := 0; i < len(text); i++ {
		dest[i] = text[i] ^ key[keyCount];
		keyCount++;
		if (keyCount >= len(key)){
			keyCount = 0;
		}
	}
	return dest;
}

func xor_decrypt(text []byte, key []byte) (dest []byte) {
	dest =  make([]byte, len(text))
	keyCount := 0;
	for i := 0; i < len(text); i++ {
		dest[i] = text[i] ^ key[keyCount];
		keyCount++;
		if (keyCount >= len(key)){
			keyCount = 0;
		}
	}
	return dest;
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
package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"
)

func part1Decode(input string) string {
	password := ""
	index := 1

	for len(password) < 8 {
		// Build string
		current := input + strconv.Itoa(index)

		hashed := md5.Sum([]byte(current))
		decoded := fmt.Sprintf("%x", hashed)

		if decoded[:5] == "00000" {
			password += string(decoded[5])
			log.Printf("Current password: %s\n", password)
		}
		index++
	}

	return password
}

func part2Decode(input string) string {
	password := []byte{'.', '.', '.', '.', '.', '.', '.', '.'}
	index := 1
	count := 0

	for count < 8 {
		// Build string
		current := input + strconv.Itoa(index)

		hashed := md5.Sum([]byte(current))
		decoded := fmt.Sprintf("%x", hashed)

		if decoded[:5] == "00000" {
			pos, _ := strconv.ParseInt(string(decoded[5]), 16, 0)
			// log.Printf("Position: %d | index: %d | decoded char: %s | hash: %x", pos, index, string(decoded[6]), hashed)
			if pos < 8 && password[pos] == '.' {
				password[pos] = decoded[6]
				log.Printf("Current password: %s\n", string(password))
				count++
			}
		}
		index++
	}

	return string(password)
}

func main() {
	input := "uqwqemis"

	password := part1Decode(input)
	fmt.Printf("Part 1: %s\n", password)

	password = part2Decode(input)
	fmt.Printf("Part 2: %s\n", password)
}

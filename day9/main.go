package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func decompress(multiple int, input string) string {
	// log.Println(input)
	s := ""
	for i := 0; i < multiple-1; i++ {
		s += input
	}
	// log.Println(s)
	return s
}

func process(input string) int {
	resultCount := 0

	skipCount := 0
	for i, c := range input {
		switch {
		case skipCount != 0: // Skip over marker
			skipCount--
		case c == '(' && skipCount == 0:
			markerRaw := ""
			for idx := i + 1; input[idx] != ')'; idx++ {
				markerRaw += string(input[idx])
			}
			marker := strings.Split(markerRaw, "x")
			numChars, _ := strconv.Atoi(marker[0])
			multiple, _ := strconv.Atoi(marker[1])

			resultCount += numChars * multiple
			skipCount += len(markerRaw) + 1 + numChars // Length of marker + ) + the actual letters to be skipped. i is already in (.
			// log.Printf("skipCount: %d | i: %d", skipCount, i)
		default:
			resultCount++
		}
	}

	return resultCount
}

func processAll(input string) int {
	resultCount := 0

	skipCount := 0
	for i, c := range input {
		switch {
		case skipCount != 0: // Skip over marker
			skipCount--
		case c == '(' && skipCount == 0:
			markerRaw := ""
			for idx := i + 1; input[idx] != ')'; idx++ {
				markerRaw += string(input[idx])
			}
			marker := strings.Split(markerRaw, "x")
			markerLen := len(markerRaw)
			numChars, _ := strconv.Atoi(marker[0])
			multiple, _ := strconv.Atoi(marker[1])

			// skipCount is markerLen, plus ), plus the number of letters to be skipped. i is already in ( so it's only + 1.
			skipCount += markerLen + 1 + numChars

			// Input is string with index (i+markerLen+2) starting from after-marker , until index + `numChars`
			thisCount := processAll(input[i+markerLen+2 : i+markerLen+2+numChars])
			for multiple != 0 {
				resultCount += thisCount
				multiple--
			}
		default:
			resultCount++
		}
	}

	return resultCount
}

func main() {
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		log.Println(input)
		count := processAll(input)
		fmt.Println(count)
	}
}

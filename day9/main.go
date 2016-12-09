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
	result := ""
	operation := ""

	inMarker := false
	nonMarkerCount := 0
	for i, c := range input {
		switch {
		case !inMarker && nonMarkerCount != 0:
			result += string(c)
			nonMarkerCount--
		case c != ')' && c != '(' && inMarker:
			operation += string(c)
		case c == '(' && !inMarker:
			inMarker = true
		case c == ')' && inMarker:
			inMarker = false
			// Perform operation
			op := strings.Split(operation, "x")
			numChars, _ := strconv.Atoi(string(op[0]))
			multiple, _ := strconv.Atoi(string(op[1]))
			result += decompress(multiple, input[i+1:i+1+numChars])
			nonMarkerCount = numChars
			operation = ""
		default:
			result += string(c)
		}
	}
	log.Println(result)

	return len(result)
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
		count := process(input)
		fmt.Println(count)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// From: https://gobyexample.com/collection-functions
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func IsABBA(input string) bool {
	for i := 0; i < len(input)-3; i++ {
		if input[i] == input[i+3] && input[i+1] == input[i+2] && input[i] != input[i+1] {
			return true
		}
	}

	return false
}

func IsXYX(input string) []string {
	abaCollection := []string{}
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] && input[i] != input[i+1] {
			byteSlice := []byte{input[i], input[i+1], input[i]}
			abaCollection = append(abaCollection, string(byteSlice))
		}
	}

	return abaCollection
}

func IsTLS(input string) bool {
	re := regexp.MustCompile("\\[([^\\]]*)\\]")

	// Parse each character
	currentLine := input
	bracketStrings := re.FindAllStringSubmatch(currentLine, -1)
	potentialStrings := re.Split(currentLine, -1)
	fmt.Printf("%s\n", currentLine)

	// nextScan := false
	// If it's bracket string is ABBA go to next scan
	for _, submatch := range bracketStrings {
		if IsABBA(submatch[1]) {
			// log.Println("Invalid hypernet")
			return false
		}
	}

	for _, submatch := range potentialStrings {
		if IsABBA(submatch) {
			// log.Println("Is ABBA")
			return true
		}
	}

	return false
}

func IsSSL(input string) bool {
	re := regexp.MustCompile("\\[([^\\]]*)\\]")

	currentLine := input
	bracketStrings := re.FindAllStringSubmatch(currentLine, -1)
	potentialStrings := re.Split(currentLine, -1)

	log.Println(currentLine)

	// Get abas
	abas := []string{}
	for _, submatch := range potentialStrings {
		res := IsXYX(submatch)
		abas = append(abas, res...)
	}
	log.Printf("abas: %q\n", abas)

	// Get babs
	babs := []string{}
	for _, submatch := range bracketStrings {
		res := IsXYX(submatch[1])
		babs = append(babs, res...)
	}
	log.Printf("babs: %q\n", babs)

	// Do comparison
	for _, aba := range abas {
		for _, bab := range babs {
			log.Printf("aba: %s | bab: %s", aba, bab)
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}

	return false
}

func main() {
	countPart1, countPart2 := 0, 0
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		// Part 1
		if IsTLS(input) {
			countPart1++
		}
		// Part 2
		if IsSSL(input) {
			countPart2++
		}

	}

	fmt.Printf("Part 1: %d\n", countPart1)
	fmt.Printf("Part 2: %d\n", countPart2)
}

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

func main() {
	count := 0
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("\\[([^\\]]*)\\]")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse each character
		// fmt.Println(scanner.Text())
		currentLine := scanner.Text()
		bracketStrings := re.FindAllStringSubmatch(currentLine, -1)
		potentialStrings := re.Split(currentLine, -1)
		fmt.Printf("%s\n", currentLine)

		nextScan := false
		// If it's bracket string is ABBA go to next scan
		for _, submatch := range bracketStrings {
			if IsABBA(submatch[1]) {
				log.Println("Invalid hypernet")
				nextScan = true
				break
			}
		}
		if nextScan {
			continue
		}

		for _, submatch := range potentialStrings {
			if IsABBA(submatch) {
				fmt.Println("Is ABBA")
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

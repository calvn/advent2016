package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var xPos, yPos, xPos2, yPos2 = 1, 1, 0, 2

	keypad := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	keypad2 := [][]int{
		[]int{0, 0, 1, 0, 0},
		[]int{0, 2, 3, 4, 0},
		[]int{5, 6, 7, 8, 9},
		[]int{0, 65, 66, 67, 0},
		[]int{0, 0, 68, 0, 0},
	}

	// Read input into slice
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))

	combos, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	password := make([]int, len(combos))
	password2 := make([]int, len(combos))

	for i, comboSlice := range combos { // Each combination
		for _, step := range comboSlice[0] { // Each step per comination
			switch step {
			case 68: // Down
				if yPos < 2 {
					yPos++
				}
				if yPos2 < 4 && keypad2[yPos2+1][xPos2] != 0 {
					yPos2++
				}
			case 85: // Up
				if yPos > 0 {
					yPos--
				}
				if yPos2 > 0 && keypad2[yPos2-1][xPos2] != 0 {
					yPos2--
				}
			case 76: // Left
				if xPos > 0 {
					xPos--
				}
				if xPos2 > 0 && keypad2[yPos2][xPos2-1] != 0 {
					xPos2--
				}
			case 82: // Right
				if xPos < 2 {
					xPos++
				}
				if xPos2 < 4 && keypad2[yPos2][xPos2+1] != 0 {
					xPos2++
				}
			}
		}
		// Record combination
		log.Printf("Part 1 - i: %d | x: %d| y: %d\n", i, xPos, yPos)
		password[i] = keypad[yPos][xPos]
		log.Printf("Part 2 - i: %d | x: %d| y: %d\n", i, xPos2, yPos2)
		password2[i] = keypad2[yPos2][xPos2]
	}

	fmt.Println(password)
	fmt.Println(password2)
}

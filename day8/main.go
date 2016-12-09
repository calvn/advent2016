package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const X = 50
const Y = 6

// Data: X = 50 Y = 6
// Example: X = 7 Y = 3
type Screen [Y][X]byte

func (s *Screen) Print() {
	for _, y := range s {
		for _, element := range y {
			fmt.Printf("%s", string(element))
		}
		fmt.Printf("\n")
	}
}

func (s *Screen) Rect(a, b int) {
	for y := 0; y < b; y++ {
		for x := 0; x < a; x++ {
			s[y][x] = '#'
		}
	}
}

func (s *Screen) RotateColumn(x, amount int) {
	// Copy original state of column into slice
	originalState := []byte{}
	for currentY := 0; currentY < Y; currentY++ {
		originalState = append(originalState, s[currentY][x])
		s[currentY][x] = '.'
	}

	// Shift slice
	for currentY := 0; currentY < Y; currentY++ {
		shiftedIndex := (currentY + amount) % Y
		s[shiftedIndex][x] = originalState[currentY]
	}
}

func (s *Screen) RotateRow(y, amount int) {
	// Copy original state of column into slice
	originalState := []byte{}
	for currentX := 0; currentX < X; currentX++ {
		originalState = append(originalState, s[y][currentX])
		s[y][currentX] = '.'
	}

	// Shift slice
	for currentX := 0; currentX < X; currentX++ {
		shiftedIndex := (currentX + amount) % X
		s[y][shiftedIndex] = originalState[currentX]
	}
}

func (s *Screen) CountLit() int {
	count := 0
	for y, row := range s {
		for x, _ := range row {
			if s[y][x] == '#' {
				count++
			}
		}
	}

	return count
}

func NewScreen() *Screen {
	screen := &Screen{}
	for y, row := range screen {
		for x, _ := range row {
			screen[y][x] = '.'
		}
	}
	return screen
}

func (s *Screen) ParseLine(input string) {
	collection := strings.Split(input, " ")
	if collection[0] == "rect" {
		xy := strings.Split(collection[1], "x")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		s.Rect(x, y)
	} else {
		pos, _ := strconv.Atoi(collection[2][2:])
		amount, _ := strconv.Atoi(collection[4])
		if collection[1] == "row" {
			s.RotateRow(pos, amount)
		} else if collection[1] == "column" {
			s.RotateColumn(pos, amount)
		}
	}

}

func main() {
	s := NewScreen()

	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		s.ParseLine(input)
	}
	s.Print()
	fmt.Printf("Part 1 - %d\n", s.CountLit())
}

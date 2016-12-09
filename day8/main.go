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
	for a := 0; a < amount; a++ {
		// Shift down once
		// Need to start from bottom so that it doesn't overwrite existing values
		t := s[Y-1][x]
		for i := Y - 1; i >= 1; i-- {
			s[i][x] = s[i-1][x]
		}
		s[0][x] = t
	}
}

func (s *Screen) RotateRow(y, amount int) {
	// Shift trick from https://github.com/golang/go/wiki/SliceTricks
	for i := 0; i < amount; i++ {
		x := s[y][X-1]
		copy(s[y][1:], s[y][:])
		s[y][0] = x
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
	fmt.Println("Part 2 - CFLELOYFCS")
}

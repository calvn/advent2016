package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Distance struct {
	Horizontal int
	Vertical   int
	Direction  Direction
}

type Block struct {
	X int
	Y int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var v = make(map[Block]bool)
var visited = false

// Part 2
func visit(x, y int) bool {
	b := Block{x, y}

	if _, ok := v[b]; !ok {
		v[b] = true
		return false
	} else {
		return true
	}
}

func (d *Distance) calculate(rawInput string, countVisits bool) error {
	// Strip whitespace
	input := strings.TrimSpace(rawInput)

	// Parse turn
	turn := string(input[0])
	if turn == "R" {
		d.Direction = (d.Direction + 1) % 4
	} else {
		d.Direction = (d.Direction - 1) % 4
		if d.Direction < 0 {
			d.Direction += 4
		}
	}

	dist, err := strconv.Atoi(string(input[1:]))
	if err != nil {
		return err
	}

	// Parse distance according to direciton
	for dist != 0 {
		switch d.Direction {
		case North:
			d.Vertical++
		case South:
			d.Vertical--
		case East:
			d.Horizontal++
		case West:
			d.Horizontal--
		}

		if visited = visit(d.Horizontal, d.Vertical); countVisits && visited {
			break
		}
		dist--
	}

	return nil
}

func main() {
	part2 := true

	// Read input into slice
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))

	records, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	d := &Distance{
		Direction: North,
	}

	for _, r := range records {
		err = d.calculate(string(r), part2)
		if err != nil {
			log.Fatal(err)
		}
		if visited && part2 {
			break
		}
	}

	distFromHQ := math.Abs(float64(d.Horizontal)) + math.Abs(float64(d.Vertical))
	fmt.Printf("X: %d | Y: %d\n", d.Horizontal, d.Vertical)
	fmt.Printf("Distance from HQ: %d\n", int64(distFromHQ))
}

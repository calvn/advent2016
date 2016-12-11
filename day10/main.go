package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Values we are looking for
const (
	MAX = 61
	MIN = 17
)

type Thing struct {
	Type  string
	Index int
}

func main() {
	// Read input into slice
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ' '
	r.FieldsPerRecord = -1

	instructions, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bots := map[int][]int{}
	bins := map[int][]int{}

	// Queue of pending bot commands, key is bot N, value is bots/bins to give to
	queue := map[int][]Thing{}

	for _, cmd := range instructions {
		// Add value to bot
		if cmd[0] == "value" {
			val, _ := strconv.Atoi(cmd[1])
			botN, _ := strconv.Atoi(cmd[len(cmd)-1])
			bots[botN] = append(bots[botN], val)
			continue
		}

		// Else bot swap instruction
		currentBot, _ := strconv.Atoi(cmd[1])
		firstThing := cmd[5]
		firstThingN, _ := strconv.Atoi(cmd[6])
		secondThing := cmd[10]
		secondThingN, _ := strconv.Atoi(cmd[11])

		thing1 := Thing{firstThing, firstThingN}
		thing2 := Thing{secondThing, secondThingN}
		queue[currentBot] = append(queue[currentBot], thing1, thing2)
	}

	for len(bots) != 0 {
		for k, v := range bots {
			if len(v) == 2 { // Parse through bots that has two chips
				sort.Ints(v)
				if v[0] == MIN && v[1] == MAX {
					fmt.Printf("Part 1 - %d\n", k)
				}

				// If no more pending commands, done with bot[k]
				if len(queue[k]) == 0 {
					continue
				}

				thing1 := queue[k][0]
				thing2 := queue[k][1]
				if thing1.Type == "bot" { // Append low to thing 1
					bots[thing1.Index] = append(bots[thing1.Index], v[0])
				} else {
					bins[thing1.Index] = append(bins[thing1.Index], v[0])
				}
				if thing2.Type == "bot" { // Append high to thing 2
					bots[thing2.Index] = append(bots[thing2.Index], v[1])
				} else {
					bins[thing2.Index] = append(bins[thing2.Index], v[1])
				}

				// Clean up, shift queue of commands from bot N
				queue[k] = queue[k][2:]
				// Remove bot[k] since chips has been distributed
				delete(bots, k)
			}
		}
	}

	part2 := bins[0][0] * bins[1][0] * bins[2][0]
	fmt.Printf("Part 2 - %d\n", part2)
}

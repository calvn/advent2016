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

var bots = map[int]*Bot{}
var bins = map[int][]int{}

type QueueCommand struct {
	First struct {
		Type  string
		Index int
	}
	Second struct {
		Type  string
		Index int
	}
}

type Bot struct {
	Id             int
	Values         []int
	PendingCommads []QueueCommand
}

// Evaluate from PendingCommads, and recursively evaluate child bots
func (b *Bot) Eval() {
	if len(b.Values) < 2 {
		return
	}

	// log.Printf("Evaluating bot: %v", b)

	sort.Ints(b.Values)
	// Check for answer
	if b.Values[0] == MIN && b.Values[1] == MAX {
		fmt.Printf("Part 1 - %d\n", b.Id)
	}

	if len(b.PendingCommads) > 0 {
		// log.Printf("%v\n", b.PendingCommads[0])
		idx1 := b.PendingCommads[0].First.Index
		idx2 := b.PendingCommads[0].Second.Index

		if _, ok := bots[idx1]; !ok {
			bots[idx1] = &Bot{}
		}
		bots[idx1].Id = idx1

		if _, ok := bots[idx2]; !ok {
			bots[idx2] = &Bot{}
		}
		bots[idx2].Id = idx2

		if b.PendingCommads[0].First.Type == "bot" {
			bots[idx1].Values = append(bots[idx1].Values, b.Values[0])
			bots[idx1].Eval()
		} else {
			bins[idx1] = append(bins[idx1], b.Values[0])
		}
		if b.PendingCommads[0].Second.Type == "bot" {
			bots[idx2].Values = append(bots[idx2].Values, b.Values[1])
			bots[idx2].Eval()
		} else {
			bins[idx2] = append(bins[idx2], b.Values[1])
		}

		// Shift commands by 1
		b.PendingCommads = b.PendingCommads[1:]

		// Cleanup bot.Values, not sure if it goes here
		b.Values = b.Values[:0]
	}
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

	for _, cmd := range instructions {
		// log.Printf("%q\n", cmd)

		// Add value to bot
		if cmd[0] == "value" {
			val, _ := strconv.Atoi(cmd[1])
			botN, _ := strconv.Atoi(cmd[len(cmd)-1])

			if _, ok := bots[botN]; !ok {
				bots[botN] = &Bot{}
			}

			bots[botN].Id = botN
			bots[botN].Values = append(bots[botN].Values, val)
			// Evaluate commands from queue when full
			bots[botN].Eval()
			continue
		}

		// Else bot swap instruction
		currentBot, _ := strconv.Atoi(cmd[1])
		firstThing := cmd[5]
		firstThingN, _ := strconv.Atoi(cmd[6])
		secondThing := cmd[10]
		secondThingN, _ := strconv.Atoi(cmd[11])

		if _, ok := bots[currentBot]; !ok {
			bots[currentBot] = &Bot{}
		}

		qc := &QueueCommand{
			First: struct {
				Type  string
				Index int
			}{
				Type:  firstThing,
				Index: firstThingN,
			},
			Second: struct {
				Type  string
				Index int
			}{
				Type:  secondThing,
				Index: secondThingN,
			},
		}

		bots[currentBot].Id = currentBot
		bots[currentBot].PendingCommads = append(bots[currentBot].PendingCommads, *qc)
		bots[currentBot].Eval()
	}

	part2 := bins[0][0] * bins[1][0] * bins[2][0]
	fmt.Printf("Part 2 - %d\n", part2)
}

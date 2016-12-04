package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Entry struct {
	Measurements [3]int
	MaxIndex     int
}

func (e *Entry) CalcMaxIndex() {
	for i, m := range e.Measurements {
		log.Printf("current: %d | max: %d | index: %d", m, e.Measurements[e.MaxIndex], i)
		if m > e.Measurements[e.MaxIndex] {
			e.MaxIndex = i
		}
	}
}

func main() {
	count := 0
	countPart2 := 0

	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ' '
	r.TrimLeadingSpace = true

	designDocuments, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range designDocuments {
		log.Printf("%+q", e)
		first, _ := strconv.Atoi(e[0])
		second, _ := strconv.Atoi(e[1])
		third, _ := strconv.Atoi(e[2])
		entry := &Entry{
			Measurements: [3]int{
				first,
				second,
				third,
			},
		}
		entry.CalcMaxIndex()
		mi := entry.MaxIndex
		log.Println(entry)

		// Determining the other two indexes: (max + 1) % 3 and (max + 2) % 3
		if entry.Measurements[(mi+1)%3]+entry.Measurements[(mi+2)%3] > entry.Measurements[mi] {
			count++
		}
	}

	log.Println("============== Part 1 Finished ==============")

	// Part 2, visiting three entries in parallel
	entries := [3]Entry{}
	for i, doc := range designDocuments {
		index := i % 3
		// Populate entries slice
		for j, m := range doc {
			val, _ := strconv.Atoi(m)
			entries[j].Measurements[index] = val
		}

		// End of third entry, perform calculation
		if (i+1)%3 == 0 {
			for _, entry := range entries {
				entry.CalcMaxIndex()
				mi := entry.MaxIndex
				log.Println(entry)

				if entry.Measurements[(mi+1)%3]+entry.Measurements[(mi+2)%3] > entry.Measurements[mi] {
					countPart2++
				}
			}
			// log.Println(entries)
		}
	}

	log.Println("============== Part 2 Finished ==============")
	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", countPart2)
}

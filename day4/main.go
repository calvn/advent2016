package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}

	return p[i].Value < p[j].Value
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

func NewRoom(entry string) *Room {
	l := len(entry)
	sectorId, _ := strconv.Atoi(entry[l-10 : l-7])

	room := &Room{
		Name:     entry[:l-11],
		SectorID: sectorId,
		Checksum: entry[l-6 : l-1],
	}

	log.Printf("%+v", room)

	return room
}

func (r *Room) IsReal() bool {
	charCount := map[rune]int{}

	// Parse encrypted name
	for _, c := range r.Name {
		if c == '-' {
			continue
		}

		// Increment char count in map
		charCount[c]++
	}

	// Sort by count and alpha
	pl := make(PairList, len(charCount))
	i := 0
	for k, v := range charCount {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	get := []byte{}

	for _, p := range pl {
		get = append(get, byte(p.Key))
	}

	// Get only first 5
	get = get[:5]

	log.Printf("Got: %s", get)

	return string(get) == r.Checksum
}

func main() {
	rooms := []Room{}
	count := 0

	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := NewRoom(scanner.Text())
		if r.IsReal() {
			count += r.SectorID
			rooms = append(rooms, *r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 - %d\n", count)

	// Part 2
	for _, r := range rooms {
		decrypted := ""

		for _, c := range r.Name {
			if c == '-' {
				decrypted += string(' ')
			} else {
				decrypted += string((int(c)-97+r.SectorID)%26 + 97) // Perform shift starting from 'a'
			}
		}
		if decrypted == "northpole object storage" {
			fmt.Printf("Part 2 - %d\n", r.SectorID)
			return
		}
	}
}

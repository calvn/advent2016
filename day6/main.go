package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func decode(charCountList []map[rune]int, leastCommon bool) string {
	secret := ""
	for _, list := range charCountList {
		pl := make(PairList, len(list))
		i := 0
		for k, v := range list {
			pl[i] = Pair{k, v}
			i++
		}
		if leastCommon {
			sort.Sort(pl)
		} else {
			sort.Sort(sort.Reverse(pl))
		}

		get := []byte{}

		for _, p := range pl {
			get = append(get, byte(p.Key))
		}

		// Get only most counted
		secret += string(get[0])
	}

	return secret
}

func NewCharCountLists() ([]map[rune]int, int) {
	return nil, 0
}

func main() {
	// Initialize slice of maps
	charCountList := []map[rune]int{}
	for i := 0; i < 8; i++ {
		charCountList = append(charCountList, map[rune]int{})
	}

	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse each character
		for i, c := range scanner.Text() {
			charCountList[i][c]++
		}
	}

	part1 := decode(charCountList, false)
	part2 := decode(charCountList, true)

	fmt.Printf("Part 1: %s\n", part1)
	fmt.Printf("Part 2: %s\n", part2)
}

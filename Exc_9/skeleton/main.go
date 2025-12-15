package main

import (
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Main function
func main() {
	// todo read file
	text := readMeditationsLines()

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)

	// todo print your result to stdout
	printTopN(results, 30)
}

func readMeditationsLines() []string {
	// First try the assignment path
	paths := []string{
		"res/meditations.txt",
		"meditations.txt",
	}

	var data []byte
	var err error

	for _, p := range paths {
		data, err = os.ReadFile(p)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("could not read meditations file : %v", err)
	}

	// Split into lines (input must be []string)
	raw := strings.Split(string(data), "\n")

	// Optional: drop empty lines
	lines := make([]string, 0, len(raw))
	for _, line := range raw {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func printTopN(freq map[string]int, n int) {
	type pair struct {
		word  string
		count int
	}
	all := make([]pair, 0, len(freq))
	for w, c := range freq {
		all = append(all, pair{word: w, count: c})
	}

	sort.Slice(all, func(i, j int) bool {
		if all[i].count == all[j].count {
			return all[i].word < all[j].word
		}
		return all[i].count > all[j].count
	})

	if n > len(all) {
		n = len(all)
	}

	fmt.Printf("Top %d words:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%4d  %s\n", all[i].count, all[i].word)
	}
}

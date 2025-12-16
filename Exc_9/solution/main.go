package main

import (
	"bufio"
	"exc9/mapred"
	"fmt"
	"os"
)

// Main function
func main() {
	// todo read file
	f, err := os.Open("res/meditations.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var text []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		text = append(text, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	for k, v := range results {
		fmt.Printf("%s: %d\n", k, v)
	}
}

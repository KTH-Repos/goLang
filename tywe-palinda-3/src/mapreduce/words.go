package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

const DataFile = "loremipsum.txt"

func spaceFinder(text string, i int) int {
	for ; i < len(text); i++ {
		if text[i] == ' ' {
			return i
		}
	}
	return i
}

func texEdit(text string, ch chan<- map[string]int) {
	text = strings.ToLower(text)
	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(text, -1)
	partitionMap := make(map[string]int)
	for _, word := range words {
		partitionMap[word]++
	}
	ch <- partitionMap
}

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	ch := make(chan map[string]int)
	freqs := make(map[string]int)
	textPartition := 16
	start := 0
	finish := spaceFinder(text, len(text)/textPartition)
	for i := 0; i < textPartition; i++ {
		if i+1 == textPartition {
			finish = len(text)
		}
		go texEdit(text[start:finish], ch)
		start, finish = finish, spaceFinder(text, finish+len(text)/textPartition)
	}

	for i := 0; i < textPartition; i++ {
		partitionMap := <-ch
		for key, value := range partitionMap {
			freqs[key] += value
		}
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, _ := ioutil.ReadFile(DataFile)
	fmt.Printf("%#v", WordCount(string(data)))
	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}

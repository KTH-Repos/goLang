// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// TODO: Answer questions.
	go questionReceiver(questions, answers)
	// TODO: Make prophecies.
	go propheciesMaker(answers)
	// TODO: Print answers.
	go answerPrinter(answers)
	return questions
}

// A function that receives all questions, and for each incoming question,
// creates a separate go-routine that answers that question
func questionReceiver(questions <-chan string, answers chan<- string) {
	for w := range questions {
		go prophecy(w, answers)
	}
}

// A function that generates predictions. It pauses for random number of minutes
// before it sends answers to the channel.
func propheciesMaker(answers chan<- string) {
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		prophecy(" ", answers)
	}
}

// A function that receives all answers and predictions, and prints them to stdout
func answerPrinter(answers <-chan string) {
	for w := range answers {
		fmt.Println(w)
		fmt.Print(prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"Roses are red.",
		"Stop barking the wrong tree.",
		"Saying I dont have time is my go to execuse",
		"It's the journey, not the destination means I have no clue where I am going",
		"Ooooh Look! Boogyman is behind you.",
		"Time heals means I'm sick of hearing about you problems",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

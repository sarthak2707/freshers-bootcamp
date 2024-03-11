package day2Exercises

import (
	"encoding/json"
	"fmt"
	"sync"
)

func incrementCount(word string, wg *sync.WaitGroup, counters *[26]chan int) {
	defer wg.Done()

	asciiA := int('a')
	for index := range word {
		go func(ind int) {
			charIndex := int(word[ind]) - asciiA
			count := <-counters[charIndex]
			count++
			counters[charIndex] <- count
		}(index)
	}
}

func getFinalCountsInMapAndCloseChannels(counters [26]chan int) map[string]int {
	finalCounts := make(map[string]int)
	asciiA := int('a')

	for i := 0; i < 26; i++ {
		lc := <-counters[i]
		if lc > 0 {
			finalCounts[string(rune(asciiA+i))] = lc
		}
		close(counters[i])
	}

	return finalCounts
}

func printMap(m map[string]int) {
	val, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(val))
	}
}

func countLetters() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	var wg sync.WaitGroup
	var counters [26]chan int

	for i := 0; i < 26; i++ {
		counters[i] = make(chan int, 1)
		counters[i] <- 0
	}

	for _, word := range words {
		wg.Add(1)
		go incrementCount(word, &wg, &counters)
	}
	wg.Wait()

	finalCounts := getFinalCountsInMapAndCloseChannels(counters)

	printMap(finalCounts)
}

func LetterCounter() {
	countLetters()
}

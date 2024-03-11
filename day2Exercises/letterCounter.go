package day2Exercises

import "sync"

func incrementCount(word string, wg *sync.WaitGroup) {
	defer wg.Done()

}

func countLetters(words []string) {
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go incrementCount(word, &wg)
	}
	wg.Wait()
}

func LetterCounter() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	countLetters(words)
}

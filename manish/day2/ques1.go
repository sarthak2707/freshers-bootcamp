package day2

import (
	"fmt"
	"sync"
)

func addFrequency(s string, wg *sync.WaitGroup, freq *[26]int) {
	defer wg.Done()
	var mutex sync.Mutex
	fmt.Println(s)
	for _, c := range s {
		mutex.Lock()
		freq[c-'a']++
		mutex.Unlock()
	}
}
func Question1() {
	inputStrings := []string{"quick", "brown", "fox", "lazy", "dog"}
	var freq [26]int

	var wg sync.WaitGroup

	for _, str := range inputStrings {
		wg.Add(1)
		//str := str
		go addFrequency(str, &wg, &freq)
	}

	wg.Wait()

	for i, f := range freq {
		fmt.Println(string(rune('a'+i)), ":", f)
	}
}

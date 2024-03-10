package day2Exercises

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func getRating(id int, rating *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Student -", id, "started rating")
	time.Sleep(time.Second * time.Duration(rand.Intn(6)))
	atomic.AddInt32(rating, rand.Int31n(6))
	fmt.Println("Student -", id, "finished rating")
}

func TeacherRating() {
	var totalRating int32
	studentCount := 200
	var wg sync.WaitGroup

	for i := 1; i <= studentCount; i++ {
		wg.Add(1)
		go getRating(i, &totalRating, &wg)
	}

	wg.Wait()

	avg := float32(totalRating) / float32(studentCount)
	fmt.Printf("Average rating for the teacher is - %f\n", avg)
}

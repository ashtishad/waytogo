package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userID := 1

	respCh := make(chan string)

	// fetch all concurrently
	go fetchUserData(userID, respCh)
	go fetchUserLikes(userID, respCh)
	go fetchUserRecommendation(userID, respCh)

	fmt.Println("time elapsed:", time.Since(start))
}

func fetchUserData(userID int, respCh chan string) {
	time.Sleep(50 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched data for user_id: %d", userID)
}

func fetchUserLikes(userID int, respCh chan string) {
	time.Sleep(50 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched likes for user_id: %d", userID)
}

func fetchUserRecommendation(userID int, respCh chan string) {
	time.Sleep(100 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched recommendations for user_id: %d", userID)
}

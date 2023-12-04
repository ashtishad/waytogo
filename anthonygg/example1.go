package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userID := 1

	// fetch all
	data := fetchUserData(userID)
	likes := fetchUserLikes(userID)
	recs := fetchUserRecommendation(userID)

	fmt.Printf("data: %s likes: %s, recos: %s\n", data, likes, recs)
	fmt.Println("time elapsed:", time.Since(start))
}

func fetchUserData(userID int) string {
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("fetched data for user_id: %d", userID)
}

func fetchUserLikes(userID int) string {
	time.Sleep(50 * time.Millisecond)
	return fmt.Sprintf("fetched likes for user_id: %d", userID)
}

func fetchUserRecommendation(userID int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("fetched recommendations for user_id: %d", userID)
}

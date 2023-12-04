package anthonygg

import (
	"fmt"
	"sync"
	"time"
)

func UserDataConcurrencyIntro() {
	start := time.Now()
	userID := 1

	var wg sync.WaitGroup

	// buffered channel is for accommodating the exact number of responses I am expecting.
	respCh := make(chan string, 3)

	wg.Add(3)

	go fetchUserData(userID, respCh, &wg)
	go fetchUserLikes(userID, respCh, &wg)
	go fetchUserRecommendation(userID, respCh, &wg)

	wg.Wait()

	close(respCh)

	for resp := range respCh {
		fmt.Println(resp)
	}

	fmt.Println("time elapsed:", time.Since(start))
}

func fetchUserData(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched data for user_id: %d", userID)
}

func fetchUserLikes(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched likes for user_id: %d", userID)
}

func fetchUserRecommendation(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)

	respCh <- fmt.Sprintf("fetched recommendations for user_id: %d", userID)
}

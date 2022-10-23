package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	// now := time.Now()
	userId := 10
	responseCh := make(chan string)

	// group that waits
	wg := &sync.WaitGroup{}
	go fetchdata(userId, responseCh,wg) 
	wg.Add(1)

	go fUR(userId, responseCh, wg)
	wg.Add(1)

	go fUL(userId,responseCh,wg)
	wg.Add(1)



	wg.Add(1)
	wg.Wait()
	close(responseCh)



	// fmt.Println(userData)
	// fmt.Println(userRex)
	// fmt.Println(userLike)
	for reps := range responseCh {
		fmt.Println(reps)
	}
}

func fetchdata (userId int , resresponseCh chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	resresponseCh <- "user data"
}

func fUR (userid int, resresponseCh chan string,wg *sync.WaitGroup){
	time.Sleep(120 * time.Millisecond)
	resresponseCh <- "user rec"
}
func fUL (userid int, resresponseCh chan string,wg *sync.WaitGroup){
	time.Sleep(50 * time.Millisecond)
	resresponseCh <- "user likes"
	 
}
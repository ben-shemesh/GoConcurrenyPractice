package main

import (
	"fmt"

)
type Message struct {
	from 	string
	payloads string
}
type Demon struct{
	messageCh chan Message
}

func (d *Demon) demonStopandListen() (){
	fmt.Println("sending message")
	for {
		msg := <- d.messageCh
		fmt.Printf("recienced message frim %s payload %s", msg.from, msg.payloads)
	}
}
func DemonServer(){

} 
func main (){
s := &Demon{
	messageCh: make(chan Message),
}
go s.demonStopandListen()

}

// func main(){
// 	// now := time.Now()
// 	userId := 10
// 	// without only one at a time
// 	// with buffer (3)
// 	responseCh := make(chan string, 3)

// 	// group that waits
// 	wg := &sync.WaitGroup{}
// 	go fetchdata(userId, responseCh,wg) 
// 	wg.Add(1)

// 	go fUR(userId, responseCh, wg)
// 	wg.Add(1)

// 	go fUL(userId,responseCh,wg)
// 	wg.Add(1)

// 	wg.Wait()
// 	close(responseCh)

// 	for reps := range responseCh {
// 		fmt.Println(reps)
// 	}
// }

// func fetchdata (userId int , resresponseCh chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	resresponseCh <- "user data"
// 	wg.Done()
// }

// func fUR (userid int, resresponseCh chan string,wg *sync.WaitGroup){
// 	time.Sleep(120 * time.Millisecond)
// 	resresponseCh <- "user rec"
// 	wg.Done()

// }
// func fUL (userid int, resresponseCh chan string,wg *sync.WaitGroup){
// 	time.Sleep(50 * time.Millisecond)
// 	resresponseCh <- "user likes"
// 	wg.Done()
	 
// }
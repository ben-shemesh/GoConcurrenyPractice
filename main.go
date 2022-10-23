package main

import (
	"fmt"
	"time"
)
type Message struct {
	from 	string
	payloads string
}
type Demon struct{
	messageCh chan Message
	quitCh chan struct{}
}

func (d *Demon) demonStopandListen(){
running:
	for {
		select{
		case msg := <- d.messageCh:
			fmt.Printf("recieved a message from %s\npayload %s\n", msg.from, msg.payloads)
		case <- d.quitCh:
			fmt.Println("the server is doing a sitdown")
			break running
		default:
		}
	}
	fmt.Println("gracefull shutdown")
}

func sendMesage(msgCh chan Message, payload string){
	msg := Message {
		from: "Joe",
		payloads: payload,
	}
	msgCh <- msg
}
func graceQuit(quitch chan struct{}) {
	close(quitch)
}
func main (){
s := &Demon{
	messageCh: make(chan Message),
	quitCh: make(chan struct{}),
}
go s.demonStopandListen()
go func (){
	time.Sleep( 2 * time.Millisecond)
	sendMesage(s.messageCh, "Yello")
}()
go func (){
	time.Sleep( 4 * time.Millisecond)
	sendMesage(s.messageCh, "Yello")
	graceQuit(s.quitCh)
}()
select{}
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
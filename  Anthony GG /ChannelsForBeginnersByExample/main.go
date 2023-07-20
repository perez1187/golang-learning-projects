package main

import (
	"fmt"
	"time"
)

type Message struct {
	From     string
	Payloads string //normaly slice of bytes
}

type Server struct {
	msgch  chan Message
	quitch chan struct{} // struct has 0 bytes
}

func (s *Server) startAndListen() {
	// so we make a loop

	// we can name loop
running:
	for {
		select {
		// we block here until someone is sending a message to the channel
		case msg := <-s.msgch:
			fmt.Printf("received message from: %s payload: %s \n", msg.From, msg.Payloads)
		case <-s.quitch:
			fmt.Println("the server is doing a gracefull shutdown")
			// logic for the gracefull shutdown
			break running //because it is a loop
		default:
		}
	}
	fmt.Println("the server us shut down")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:     "Joe",
		Payloads: payload,
	}

	msgch <- msg
}

func graceFullQuit(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	// so we say: go and listen
	go s.startAndListen()

	// send a messsage in another routine

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "yo yo yo yo")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		graceFullQuit(s.quitch)
	}()

	// we block program and testing. Never do it in prod
	select {}
}

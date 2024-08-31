package main

import (
	"fmt"
	"time"
)

func greet(phrase string, daneChan chan bool) {
	fmt.Println("Hello!", phrase)
	daneChan <- true
}

func slowGreet(phrase string, daneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	daneChan <- true
}

func main() {
	dane := make(chan bool)

	go greet("Nice to meet you!", dane)
	go greet("How are you?", dane)
	go slowGreet("How ... are ... you ...?", dane)
	go greet("I hope you're liking the course!", dane)
	<-dane
	<-dane
	<-dane
	<-dane
}

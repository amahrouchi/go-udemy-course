package main

import (
	"fmt"
	"time"
)

func client1(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Message from client 1 => %d", i)
		time.Sleep(1500 * time.Millisecond)
	}
}

func client2(c chan string) {
	for i := 10; i < 15; i++ {
		c <- fmt.Sprintf("Message from client 2 => %d", i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go client1(c1)
	go client2(c2)

	maxEmpty := 10
	currEmpty := 0
	for currEmpty <= maxEmpty {
		time.Sleep(time.Second)
		select {
		case v := <-c1:
			fmt.Println("Received from client 1: ", v)
		case v := <-c2:
			fmt.Println("Received from client 2: ", v)
		default:
			// Si rien n'est prêt à être lu dans aucun channel alors le default est exécuter
			// Grace au "default" les lectures deviennent donc non bloquante
			// Sans "default" le channel sera donc bloquant et des deadlocks peuvent arriver si aucun channel ne répond

			fmt.Println("No value received")
			currEmpty++
		}

	}
}

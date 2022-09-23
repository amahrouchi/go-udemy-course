package main

import (
	"fmt"
	"time"
)

// Channels: canaux faisant communiquer des goroutines entre elles
// les channels sont bloquant l'écriture ne peut se faire que si qqun est prêt à lire
// et inversement, on ne peut lire que si qqun est prêt à écrire dans le channel
// Celui qui écrit est limité par la vitesse de lecture de celui qui écoute le channel
// Celui qui lit ne peut lire que quand un message est disponible
// Dans les 2 cas, l'exécution (R ou W) est bloqué en attendant la disponibilité du l'autre cote du channel

func ping(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func pong(c chan string) {
	for i := 100; ; i += 100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string) {
	for {
		message := <-c
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	time.Sleep(10 * time.Second)
}

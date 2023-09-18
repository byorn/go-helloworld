package main

import (
	"fmt"
	"net/http"
	"time"
)

func testGoRoutinesAndChannels() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
	}

	//create a brand new channel called c
	//if will receive values from go routines
	c := make(chan string)

	for _, link := range links {
		// by having go in front as a keyword, it means run this function
		// in a brand new go routine (like a new thread)
		// which means if the checkLink function gets blocked, the line of code
		// in the above for loop will execute.

		//its like a thread, but using one CPU. so its not 100% parallel, cause the CPU
		//will run one thread at a time.

		// concurrency is not parallelism.
		// this is concurrent, but it doesnt work parallely, cause by default its only one cpu that
		// is spawniing a new thread
		go checkLink(link, c)
	}

	//time.Sleep(10 * time.Second)
	//WILL WAIT UNTIL C GETS A MESSAGES.
	//println(<-c)

	// will wait len(links)-1 times before exiting.
	for i := 0; i < len(links); i++ {
		//go func() {
		time.Sleep(3 * time.Second)
		//receive the value from the channel
		fmt.Println(<-c)
		//}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "error"
		return
	}

	// send the link value to the channel
	c <- link
	fmt.Println(link, "go routing is exiting")
}

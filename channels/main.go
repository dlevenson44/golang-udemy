package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// new channel
	c := make(chan string)

	// below code runs through links array in order without the `go` keyword
	// can cause delay because this is running in a single-thread (sync)
	for _, link := range links {
		// checkLink((link))
		go checkLink(link, c)
	}

	// receive value from channel
	// waiting on message to come from channel is blocking line of code
	// prints out first message and then exits because main thinks it's done
	// fmt.Println((<-c))
	// fmt.Println((<-c))

	// for i := 0; i < len(links); i++ {
	// for {
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
			// passing l as argument ensures it looks at proper channel l value
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // <-- blocking call, code cannot continue until this function completes, can fix with multiple routines
	if err != nil {
		fmt.Println(link, "might be down!")
		// value to send into channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

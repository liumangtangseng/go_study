package main

import "fmt"

func main() {
	ci := make(chan int, 2)
	for i := 1; i <= 2; i++ {
		ci <- i
	}
	close(ci)
	cs := make(chan string, 2)
	cs <- "hi"
	cs <- "golang"
	close(cs)
	ciClosed, csClosed := false, false
	for {
		if ciClosed && csClosed {
			return
		}
		select {
		case i, ok := <-ci:
			if ok {
				fmt.Println(i)
			} else {
				ciClosed = true
				fmt.Println("ciclosed")
			}
		case s, ok := <-cs:
			if ok {
				fmt.Println(s)
			} else {
				csClosed = true
				fmt.Println("csclosed")
			}
		default:
			fmt.Println("waiting...")
		}
	}
}

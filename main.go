package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch := make(chan struct{})

	go f(ch)

	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(time.Second)

	ch2 := make(chan func())
	ch2 <- processTask
	go timeout(ch2)

	time.Sleep(3 * time.Second)
	close(ch2)
	time.Sleep(time.Second)


	// Graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// run main process/jobs/etc ...

	<-signalCh
	// stop ...
}

func processTask() {
	fmt.Println("do task")
}

func timeout(taskCh chan func()) {
	timer := time.NewTimer(time.Second * 3)
	select {
	case task := <- taskCh:
		task()
	case <-timer.C:
		fmt.Println("timed out!")
	}
}

func f(stopCh chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	msgCh := make(chan string)
	msgCh <- "Hi there!"
	for {
		select {
		case <-stopCh: // Priority pattern
			return
		default:
		}

		select {
		case <-stopCh:
			return
		case <-ticker.C:
			fmt.Println("task done")
		case msg := <-msgCh:
			processMsg(msg)
		}
	}
}

func processMsg(msg string) {
	fmt.Println(msg)
}
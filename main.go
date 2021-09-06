package main

import (
	"log"
	"time"
)

var BlockWriteTime = []int64{1, 11, 40, 51}

type blockChain struct{}

func (b blockChain) first() {
	//time.Sleep(time.Second * 15)
	log.Println("first")
}

func (b blockChain) second() {
	log.Println("second")
}

func (b blockChain) third() {
	log.Println("third")
}

func (b blockChain) fourth() {
	log.Println("fourth")
}

func main() {
	var b blockChain

	queue := make(chan int64)
	i := getQueue()
	go func() {
		for {
			if i != getQueue() {
				queue <- getQueue()
				i = getQueue()
			}
		}
	}()

	for {
		select {
		case j := <-queue:
			switch j {
			case BlockWriteTime[0]:
				go b.first()
				break
			case BlockWriteTime[1]:
				go b.second()
				break
			case BlockWriteTime[2]:
				go b.third()
				break
			case BlockWriteTime[3]:
				go b.fourth()
				break
			}
		}
	}

}

func TimestampUnix() int64 {
	timestamp, err := time.Parse(time.RFC3339Nano, time.Now().Format(time.RFC3339Nano))
	if err != nil {
		log.Println("Timestamp Unix error:", err)
	}

	return timestamp.UnixNano()
}

func getQueue() int64 {
	return ((TimestampUnix() / time.Second.Nanoseconds()) % BlockWriteTime[3]) + 1
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	startDate := time.Now()
	// add 10 seconds to current time
	startDate = startDate.Add(6 * time.Second)

	timer := time.NewTimer(startDate.Sub(time.Now()))
	//timer := time.NewTimer(10 * time.Second)

	go func() {
		defer timeTrack(time.Now(), "comienzo de elección")
		<-timer.C
		fmt.Println("Comenzó la elección")
	}()

	// read character from stdin

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	print(text)

	// var input string
	// done := make(chan bool)
	// print(input)
	// if input == "exit" {
	// 	done <- true
	// }
	// <-done
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
package main

import (
	"fmt"
	"log"
	"time"
)

// time tracker
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function: %s took %s\n", name, elapsed)
}

func main() {
	fileName := "./data/test-movie.mp4"

	// define chunks as 5MiB (5,242,880 in bytes)
	chunckSize := 5242880

	buff := byteReader(fileName, chunckSize)

	byteWriter("./data/test.mp4", buff)

	fmt.Printf("%d bytes written to disk\n", len(buff))
}

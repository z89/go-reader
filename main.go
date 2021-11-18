// Copyright 2017 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// time tracker
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function: %s took %s\n", name, elapsed)
}

func byteReader(path string, chunkSize int) []byte {
	defer timeTrack(time.Now(), "byteReader")
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	bytes, chunks := int(0), int(0)

	reader := bufio.NewReader(file)

	temp := make([]byte, 0, chunkSize)         // 5MiB byte slice with 0 elements (initalise)
	buffer := make([]byte, 0, fileInfo.Size()) // 5MiB byte slice with 0 elements (initalise)

	for {
		n, err := reader.Read(temp[:cap(temp)])
		temp = temp[:n]

		buffer = append(buffer, temp...)

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			panic(err)
		}
		chunks++
		fmt.Printf("bytes added: %s\n", strconv.Itoa(int(bytes)))
		fmt.Printf("chunk: %s\n", strconv.Itoa(int(chunks)))

		bytes += int(len(buffer))
		if err != nil && err != io.EOF {
			panic(err)
		}
	}

	return buffer
}

func byteWriter(path string, buffer []byte) {
	defer timeTrack(time.Now(), "byteWriter")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("writing buffer to file...\n")
	if _, err = file.Write([]byte(buffer)); err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
}

func main() {
	fileName := "./data/test-movie.mp4"

	// define chunks as 5MiB (5,242,880 in bytes)
	chunckSize := 5242880

	buff := byteReader(fileName, chunckSize)

	byteWriter("./data/test.mp4", buff)

	fmt.Printf("%d bytes written to disk\n", len(buff))
}

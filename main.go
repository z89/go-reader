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

// took 335.096889ms to read a 2.9GB movie file

func main() {
	defer timeTrack(time.Now(), "main")

	chunckSize := 5242880 // 5MiB chunks (in bytes)

	fileName := "./test-movie.mp4"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	bytes, chunks := int(0), int(0)

	reader := bufio.NewReader(file)

	buffer := make([]byte, 0, chunckSize)          // 5MiB byte slice with 0 elements (initalise)
	fullBuffer := make([]byte, 0, fileInfo.Size()) // 5MiB byte slice with 0 elements (initalise)

	newFile, err := os.OpenFile("rewrite.mp4", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	for {
		n, err := reader.Read(buffer[:cap(buffer)])
		buffer = buffer[:n]

		fullBuffer = append(fullBuffer, buffer...)

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

	fmt.Printf("writing buffer to file...\n")
	if _, err = newFile.Write([]byte(fullBuffer)); err != nil {
		panic(err)
	}

	defer func() {
		if err := newFile.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("data len: %d bytes written to disk\n", len(fullBuffer))

}

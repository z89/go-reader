package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

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

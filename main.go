package main

import (
	"context"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/semaphore"
	"knockknocker/requester"
	"os"
	"strings"
	"sync"
)

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	return lines
}

type invalidUrl struct {
	url   string
	error error
}

func main() {
	readFilePath := os.Args[1]
	writeFilePath := os.Args[2]

	invalidUrlsChain := make(chan invalidUrl)
	lines := readLines(readFilePath)
	wg := sync.WaitGroup{}
	wg.Add(len(lines))
	completeChain := make(chan bool)
	go func() {
		ctx := context.TODO()
		sem := semaphore.NewWeighted(400)
		for _, line := range lines {
			line := line
			sem.Acquire(ctx, 1)
			go func() {
				defer sem.Release(1)
				defer wg.Done()
				parsedUrl := strings.Split(line, ",")
				accessError := requester.TouchWebsite(parsedUrl[1])
				if accessError != nil {
					invalidUrlsChain <- invalidUrl{url: parsedUrl[1], error: accessError}
				}
				completeChain <- true
			}()
		}
		wg.Wait()
		close(completeChain)
		close(invalidUrlsChain)
	}()

	go func() {
		writeFile, err := os.OpenFile(writeFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		for invalidUrl := range invalidUrlsChain {
			n, err := writeFile.Write([]byte(invalidUrl.url + " " + invalidUrl.error.Error() + "\n\n\n"))
			if err != nil {
				fmt.Println(n, err)
			}
		}

		writeFileErr := writeFile.Close()
		if writeFileErr != nil {
			fmt.Println(writeFileErr)
		}

	}()

	bar := progressbar.Default(int64(len(lines)))
	for range completeChain {
		bar.Add(1)
	}
}

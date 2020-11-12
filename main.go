package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// AddToFile works
func AddToFile(s string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	str := strings.Split(s, ",")
	ch <- str[0] + " - " + str[1]
}

func runChannel(wg *sync.WaitGroup, ch chan string) {
	defer close(ch)

	wg.Wait()
}

func main() {
	start := time.Now()
	file, err := os.Open("1000000.csv")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	ch := make(chan string)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var wg sync.WaitGroup

	counter := 0

	for scanner.Scan() {
		wg.Add(1)
		go AddToFile(scanner.Text(), ch, &wg)
	}

	go runChannel(&wg, ch)

	for range ch {
		counter++
	}

	fmt.Println(counter)

	file.Close()

	fmt.Printf("Time: %s", time.Since(start))
}

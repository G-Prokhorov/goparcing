package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"strconv"
)

var num1, num2 int

// AddToFile works
func AddToFile(s string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	str := strings.Split(s, ",")
	ch <- str[num1] + " - " + str[num2]
}

func runChannel(wg *sync.WaitGroup, ch chan string) {
	defer close(ch)

	wg.Wait()
}

func main() {
	start := time.Now()

	if len(os.Args) < 4 {
		log.Fatalf("Missed argument");
	}

	a, err1 := strconv.Atoi(os.Args[2])
	b, err2 := strconv.Atoi(os.Args[3])

	num1 = a
	num2 = b

	if err1 != nil || err2 != nil {
		log.Fatalf("third and fourth element is not a number");
	}

	file, err := os.Open(os.Args[1])

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

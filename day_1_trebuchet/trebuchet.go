package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	lines := make(chan string, runtime.NumCPU())
	calibrationValues := make(chan int, 10000)
	var wg sync.WaitGroup
	var waitForSum sync.WaitGroup
	go readCalibrationDocument(lines)
	for line := range lines {
		if line == "\n" {
			break
		}
		wg.Add(1)
		go calcCalibrationValue(line, calibrationValues, &wg)
	}
	waitForSum.Add(1)
	go collectValues(calibrationValues, &waitForSum)
	wg.Wait()
	close(calibrationValues)
	waitForSum.Wait()
}

func readCalibrationDocument(ch chan<- string) {
	file, err := os.Open("day_1_trebuchet/calibration_document.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(file)
	defer close(ch)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calcCalibrationValue(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	characters := regexp.MustCompile(`[a-zA-Z]`)
	onlyDigits := characters.ReplaceAllString(line, "")
	num, err := strconv.Atoi(fmt.Sprintf("%c%c", onlyDigits[0], onlyDigits[len(onlyDigits)-1]))
	if err != nil {
		panic(err)
	}
	ch <- num
}

func collectValues(ch <-chan int, wg *sync.WaitGroup) {
	sum := 0
	for calibrationValue := range ch {
		sum += calibrationValue
	}
	fmt.Println(sum)
	wg.Done()
}

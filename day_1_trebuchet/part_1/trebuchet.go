package main

import (
	"advent-of-code-golang/common"
	"fmt"
	_ "net/http/pprof"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) == 3 && os.Args[2] == "worker" {
		fmt.Println(runAsWorkerPool(os.Args[1]))
	} else {
		fmt.Println(run(os.Args[1]))
	}
}

func run(path string) int {
	lines := make(chan string, runtime.NumCPU())
	calibrationValues := make(chan int, 10000)
	result := make(chan int)
	var wg sync.WaitGroup
	go readCalibrationDocument(lines, path)
	for line := range lines {
		if line == "\n" {
			break
		}
		wg.Add(1)
		go calcCalibrationValue(line, calibrationValues, &wg)
	}
	go collectValues(calibrationValues, result)
	wg.Wait()
	close(calibrationValues)
	return <-result
}

func runAsWorkerPool(path string) int {
	lines := make(chan string, 10000)
	calibrationValues := make(chan int, 10000)
	result := make(chan int)
	var wg sync.WaitGroup
	go readCalibrationDocument(lines, path)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go calcCalibrationValueWorker(lines, calibrationValues, &wg)
	}
	go collectValues(calibrationValues, result)
	wg.Wait()
	close(calibrationValues)
	return <-result
}

func readCalibrationDocument(ch chan<- string, path string) {
	common.ReadFile(path, func(s string) {
		ch <- s
	})
	close(ch)
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

func calcCalibrationValueWorker(lines <-chan string, ch chan<- int, wg *sync.WaitGroup) {
	for line := range lines {
		characters := regexp.MustCompile(`[a-zA-Z]`)
		onlyDigits := characters.ReplaceAllString(line, "")
		num, err := strconv.Atoi(fmt.Sprintf("%c%c", onlyDigits[0], onlyDigits[len(onlyDigits)-1]))
		if err != nil {
			panic(err)
		}
		ch <- num
	}
	wg.Done()
}

func collectValues(ch <-chan int, result chan<- int) {
	sum := 0
	for calibrationValue := range ch {
		sum += calibrationValue
	}
	result <- sum
}

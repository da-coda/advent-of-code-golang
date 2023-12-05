package common

import (
	"bufio"
	"log"
	"log/slog"
	"os"
)

func ReadFile(path string, f func(string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

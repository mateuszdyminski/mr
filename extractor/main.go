package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	dat, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	newFIle, err := os.Create("extracted")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bufio.NewReader(dat))

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		newFIle.Write([]byte(parts[2] + "\n"))
	}

	newFIle.Close()
}

package main

import (
	"container/list"
	"fmt"
	"os"
	"strings"

	"github.com/mateuszdyminski/mr/mapreduce"
)

func Map(value string) *list.List {
	list := list.New()

	lines := strings.Split(value, "\n")

	for _, l := range lines {
		// ommit empty lines
		if l == "" {
			continue
		}

		numbers := strings.Split(strings.TrimSpace(l), ",")

		for _, n := range numbers {
			list.PushFront(mapreduce.KeyValue{Key: n, Value: ""})
		}
	}

	return list
}

func Reduce(key string, values *list.List) string {
	counter := 0
	for e := values.Front(); e != nil; e = e.Next() {
		counter++
	}

	return fmt.Sprintf("%d", counter)
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master x.txt sequential)
// 2) Master (e.g., go run wc.go master x.txt localhost:7777)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	if len(os.Args) != 4 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	} else if os.Args[1] == "master" {
		if os.Args[3] == "sequential" {
			mapreduce.RunSingle(5, 1, os.Args[2], Map, Reduce)
		} else {
			mr := mapreduce.MakeMapReduce(5, 1, os.Args[2], os.Args[3])
			// Wait until MR is done
			<-mr.DoneChannel
		}
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], Map, Reduce, 100)
	}
}

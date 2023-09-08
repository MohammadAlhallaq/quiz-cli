package main

import (
	"flag"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in a format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		println(err.Error())
	}
	print(file)
}

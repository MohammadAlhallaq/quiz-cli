package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in a format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(err.Error())
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(err.Error())
	}

	problems := parseLines(lines)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			exit(err.Error())
		}
		if answer == p.a {
			fmt.Println("Correct!")
		}
	}

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			line[0],
			line[1],
		}
	}
	return ret
}

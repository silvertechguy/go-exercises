package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const defaultProblemsFilename = "problems.csv"

func main() {
	f, err := os.Open(defaultProblemsFilename)
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed to read csv file: %v\n", err)
		return
	}

	var correctAnswers int
	for index, record := range questions {
		question, correctAnswer := record[0], record[1]
		fmt.Printf("%d. %s?\n", index+1, question)
		var answer string
		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("failed to scan: %v\n", err)
			return
		}
		if answer == correctAnswer {
			correctAnswers++
		}
	}

	fmt.Printf("Result: %d / %d\n", correctAnswers, len(questions))

}

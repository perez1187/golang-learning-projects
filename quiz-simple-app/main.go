package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// read all the problems from quiz.csv

	// ( if else is a good practice to handle errors)

	// 1. open the file
	if fObj, err := os.Open(fileName); err == nil { // two conditions in if

		// 2. we will create a new reader
		csvR := csv.NewReader(fObj)

		// 3. it will need to read the file
		if cLines, err := csvR.ReadAll(); err == nil {
			// 4. call the parse Problem func -> send all file to parseProblem
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv"+"format from from %s file, %s", fileName, err.Error()) // and we print out the error
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file, %s", fileName, err.Error())
	}

}

func main() {
	// 1. input the name of the file
	fileName := flag.String("f", "quiz.csv", "path of csv file")

	// 2. set the duration of solve the quiz
	timer := flag.Int("t", 30, "timer for the quiz")

	flag.Parse() // we need to parse flags

	// 3. pull the problems from the file (calling problem puller func)
	problems, err := problemPuller(*fileName) // so if we use variable in func we need to write *nameOfVar (?)

	// 4 handle the error
	if err != nil {
		exit(fmt.Sprintf("something went wrong: %s\n", err.Error()))
	}

	// 5. create a variable to count correct answers
	correctAnswers := 0

	// 6. initial  timer
	tObject := time.NewTimer(time.Duration(*timer) * time.Second) // this is how we set timer
	answersChannel := make(chan string)

	// 7. loop through the problems, pront question, accept answer
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.question) // so p is a problem, i+1, because i start from 0

		go func() {
			fmt.Scanf("%s", &answer)
			answersChannel <- answer
		}() // this () close
		select {
		// two things can happen, you answer all question, timeout
		case <-tObject.C:
			fmt.Println()
			break problemLoop
		case iAns := <-answersChannel:
			if iAns == p.question {
				correctAnswers++
			}
			if i == len(problems)-1 {
				// close this channel
				close(answersChannel)
			}
		}
		// 8. calculate and print results
		fmt.Printf("Your result is %d out of %d\n", correctAnswers, len(problems))
		fmt.Printf("Press enter to exit")
		// enter value
		<-answersChannel
	}

}

// lines - rows from csv file
func parseProblem(lines [][]string) []problem { //[][] slice of slices? or because of two indexes?
	// go over the lines and parse them, with problem struct
	r := make([]problem, len(lines)) // []problem -> slice of problem

	// we go for every single line
	for i := 0; i < len(lines); i++ {
		r[i] = problem{question: lines[i][0], answer: lines[i][1]}
	}
	return r
}

type problem struct {
	question string
	answer   string
}

// function that exit the program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // exit with status 1
}

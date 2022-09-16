package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

type userScore struct {
	correctAnswerCount   int
	incorrectAnswerCount int
}

var wg sync.WaitGroup // create a wait group

func allArgParser() string {
	var filePath *string = flag.String(
		"filepath",
		"./problems.csv",
		"provide filepath of the csv here",
	)

	flag.Parse()
	fmt.Println("filePath has value", *filePath)
	return *filePath
}

func customReader(ctx context.Context) (int, error) {
	var x int
	fmt.Scan(&x)
	return x, errors.New("error while reading")
}

func questionConsumer(questionChannel chan []string) userScore {
	var userAnswer int = 0
	var user_score userScore
	for x := range questionChannel {
		variable, err := questionCreatorFromString(x[0])
		if err != nil {
			log.Panic("shit once")
		}
		correctAnswer, err := strconv.Atoi(x[1])
		if err != nil {
			log.Panic("shit twice")
		}
		fmt.Printf("what is %d + %d\n", variable[0], variable[1])
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(1 * time.Second)
			if userAnswer == 0 {
				fmt.Println("1 sec ended")
				cancel()
			}
		}()
		userAnswer, err = customReader(ctx)
		// wait for 30 sec and if input not provided then cancel this function which will calculate total correct for the user

		if userAnswer == correctAnswer {
			user_score.correctAnswerCount += 1
			fmt.Println("you rock")
		} else {
			user_score.incorrectAnswerCount += 1
			fmt.Println("wrong answer")
		}

	}
	return user_score
}

func questionCreatorFromString(str string) ([]int, error) {
	variables := strings.Split(str, "+")
	var1, err := strconv.Atoi(variables[0])
	if err != nil {
		return nil, err
	}
	var2, err := strconv.Atoi(variables[1])
	if err != nil {
		return nil, err
	}
	return []int{var1, var2}, nil
}

func questionProducer(filePath string, questionChannel chan []string) {
	records := readCsvFile(filePath)
	defer wg.Done()
	for _, item := range records {
		questionChannel <- item
	}
	close(questionChannel)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

func main() {
	filePath := allArgParser()
	questionChannel := make(chan []string)
	wg.Add(1)
	go questionProducer(filePath, questionChannel)
	user_score := questionConsumer(questionChannel)
	wg.Wait()
	fmt.Printf("you got %v correct and %v incorrect", user_score.correctAnswerCount, user_score.incorrectAnswerCount)

}

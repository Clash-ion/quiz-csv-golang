package main

import (
	"context"
	"fmt"

	"github.com/Clash-ion/quiz-csv-golang/pkg/arg_parser"
	"github.com/Clash-ion/quiz-csv-golang/pkg/questions"
	"github.com/Clash-ion/quiz-csv-golang/pkg/utils"
)

func main() {
	filePath, timeBetweenQuestion := arg_parser.AllArgParser()
	fmt.Println(timeBetweenQuestion)
	var userScore int
	ctx := context.Background()
	questionChannel := make(chan []string)
	questionConsumerChannel := make(chan int)
	cancelOrder := make(chan bool)
	utils.Wg.Add(1)
	go questions.QuestionProducer(filePath, questionChannel, cancelOrder, ctx)
	utils.Wg.Add(1)
	go questions.QuestionConsumer(questionChannel, questionConsumerChannel, cancelOrder, ctx)
	utils.Wg.Add(1)
	go func() {
		for score := range questionConsumerChannel {
			userScore += score
		}
		if !questions.Completed {
			defer utils.Wg.Done()
		}
		// defer utils.Wg.Done()
	}()
	utils.Wg.Wait()
	fmt.Printf("you got %v correct out of %v ", userScore, questions.TotalQuestion)

}

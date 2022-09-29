package questions

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// "github.com/Clash-ion/quiz-csv-golang/pkg/arg_parser"
	"github.com/Clash-ion/quiz-csv-golang/pkg/utils"
)

var TotalQuestion int = 0

// type UserScore struct {
// 	CorrectAnswerCount   int
// 	IncorrectAnswerCount int
// }

func QuestionConsumer(questionChannel chan []string, questionConsumerChannel chan int, cancelOrder chan bool, ctx context.Context) {
	var userAnswer int = 0
	// var user_score UserScore
	// fmt.Println()
	defer utils.Wg.Done()

	for x := range questionChannel {
		TotalQuestion += 1
		variable, err := questionCreatorFromString(x[0])
		if err != nil {
			log.Panic("shit once")
		}
		correctAnswer, err := strconv.Atoi(x[1])
		if err != nil {
			log.Panic("shit twice")
		}
		fmt.Printf("what is %d + %d\n", variable[0], variable[1])
		// _, cancel := context.WithCancel(ctx)
		// go func() {
		// 	time.Sleep(3 * time.Second)
		// 	if userAnswer == 0 {
		// 		fmt.Println("3 sec ended")
		// 		cancel()
		// 		// return
		// 	}
		// }()
		// ch := make(chan int)
		// err = utils.CustomReader(ch)
		// userAnswer = <-ch
		timeInstA := time.Now()
		fmt.Scan(&userAnswer)
		timeInstB := time.Now()
		if timeInstB.Sub(timeInstA).Seconds() > 3 {
			// cancel()
			// close(questionConsumerChannel)
			for i := 0; i < 3; i++ {
				utils.Wg.Done()
			}

		}
		// wait for 3 sec and if input not provided then cancel this function which will calculate total correct for the user
		// if err != nil {
		// 	log.Panic(err)
		// }
		if userAnswer == correctAnswer {
			questionConsumerChannel <- 1
			fmt.Println("you rock")
		} else {
			// user_score.IncorrectAnswerCount += 1
			fmt.Println("wrong answer")
		}

	}
	// abcd
	close(questionConsumerChannel)
	fmt.Println("")
	// return user_score
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

func QuestionProducer(filePath string, questionChannel chan []string, cancelOrder chan bool, ctx context.Context) {
	records := utils.ReadCsvFile(filePath)
	defer utils.Wg.Done()
	// cancel := context.CancelFunc
	_, cancel := context.WithCancel(ctx)
	// cancel()
	defer cancel()
	for _, item := range records {
		// wg.Done()

		questionChannel <- item

	}
	close(questionChannel)
}

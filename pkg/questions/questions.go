package questions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Clash-ion/quiz-csv-golang/pkg/utils"
)

var TotalQuestion int = 0

var Completed bool = false

func QuestionConsumer(questionChannel chan []string, questionConsumerChannel chan int, timeOutTime int) {
	var userAnswer int = 0
	for x := range questionChannel {
		if !Completed {
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

			timeInstA := time.Now()
			fmt.Scan(&userAnswer)
			timeInstB := time.Now()
			if err != nil {
				log.Fatal(err)
			}
			if timeInstB.Sub(timeInstA).Seconds() > float64(timeOutTime) {
				Completed = true
				fmt.Println("timeout")
				for i := 0; i < 3; i++ {
					utils.Wg.Done()
				}
			}
			if !Completed {
				if userAnswer == correctAnswer {
					questionConsumerChannel <- 1
					fmt.Println("you rock")
				} else {
					fmt.Println("wrong answer")
				}
			}
		}
	}
	close(questionConsumerChannel)
	if !Completed {
		defer utils.Wg.Done()
	}

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

func QuestionProducer(filePath string, questionChannel chan []string) {
	records := utils.ReadCsvFile(filePath)

	for _, item := range records {
		questionChannel <- item
	}
	close(questionChannel)

	if !Completed {
		defer utils.Wg.Done()
	}
}

package arg_parser

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func AllArgParser() (string, int) {
	var filePath *string = flag.String(
		"filepath",
		"./problems.csv",
		"provide filepath of the csv here",
	)
	var timeBetweenQuestion *int = flag.Int(
		"timebetweenquestion",
		3,
		"provide filepath of the csv here",
	)

	flag.Parse()
	fmt.Println("filePath has value", *filePath)
	return *filePath, *timeBetweenQuestion
}

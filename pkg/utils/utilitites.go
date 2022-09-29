package utils

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
)

var Wg sync.WaitGroup // create a wait group

func ReadCsvFile(filePath string) [][]string {
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

// func CustomReader(ch chan int) error {
// 	var x int
// 	_, err := fmt.Scan(&x)
// 	if err != nil {
// 		return err
// 	}
// 	ch <- x
// 	return nil
// }

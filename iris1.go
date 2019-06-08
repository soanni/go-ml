package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("iris1.data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5

	// rawCSVData, err := reader.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, line := range rawCSVData {
	// 	for _, field := range line {
	// 		fmt.Print(field, " ")
	// 	}
	// 	fmt.Print("\n")
	// }
	var rawCSVData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			continue
		}
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}

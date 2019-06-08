package main

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	sepalLength := irisDF.Col("sepal_length").Float()
	sepalWidth := irisDF.Col("sepal_width").Float()
	cor, _ := stats.Correlation(sepalLength, sepalWidth)
	fmt.Println(cor)
}

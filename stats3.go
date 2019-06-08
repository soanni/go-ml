package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/gonum/floats"
	"github.com/gonum/stat"
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

	petalLength := irisDF.Col("petal_length").Float()

	meanVal := stat.Mean(petalLength, nil)
	modeVal, modeCount := stat.Mode(petalLength, nil)
	medianVal, err := stats.Median(petalLength)
	if err != nil {
		log.Fatal(err)
	}

	// dispersion measures

	minVal := floats.Min(petalLength)
	maxVal := floats.Max(petalLength)
	rangeVal := maxVal - minVal

	varianceVal := stat.Variance(petalLength, nil)
	stdDevVal := stat.StdDev(petalLength, nil)

	inds := make([]int, len(petalLength))
	floats.Argsort(petalLength, inds)

	quant25 := stat.Quantile(0.25, stat.Empirical, petalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, petalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, petalLength, nil)

	fmt.Printf("\nPetal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n", medianVal)
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n", quant75)
}

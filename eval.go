package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gonum/stat"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("continuous.data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed, predicted []float64

	line := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if line == 1 {
			line++
			continue
		}

		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// mAE - mean absolute error
	// mSE - mean squared error
	var mAE, mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// R^2 value
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n", mSE)
	fmt.Printf("\nR^2 = %0.2f\n\n", rSquared)
}

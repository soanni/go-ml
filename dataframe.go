package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	f, err := os.Open("iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	irisDF := dataframe.ReadCSV(f)
	// fmt.Println(irisDF)
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)
	fmt.Println(versicolorDF.Select([]string{"sepal_width", "species"}))
	fmt.Println(versicolorDF.Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2}))
}

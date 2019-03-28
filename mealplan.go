package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Input is an input in a recipe (basically an ingredient)
type Input struct {
	name          string
	category      string
	calories      int
	quantity      int
	unit          string
	dryMultiplier float32
}

type meal struct {
	base    string
	filling string
	style   string
}

var mealCount int
var bases []Input
var fillings []Input

func init() {
	flag.IntVar(&mealCount, "m", 10, "number of meals for the plan")
	// load the data and stuff
	bases = loadData(bases, "data/bases.csv")
	fillings = loadData(fillings, "data/fillings.csv")
	// fmt.Println(fillings)
}

func main() {
	flag.Parse()
	generateMeals()
}

func generateMeals() {
	fmt.Println(fmt.Sprintf("gonna generate %d meals", mealCount))
	var meals []meal
	fmt.Println(meals)
}

func loadData(arr []Input, fileName string) []Input {
	inputFile, _ := os.Open(fileName)
	r := csv.NewReader(bufio.NewReader(inputFile))
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		calories, _ := strconv.Atoi(line[2])
		quantity, _ := strconv.Atoi(line[3])
		dryMultiplier, _ := strconv.ParseFloat(line[5], 32)
		arr = append(arr, Input{
			name:          line[0],
			category:      line[1],
			calories:      calories,
			quantity:      quantity,
			unit:          line[4],
			dryMultiplier: float32(dryMultiplier),
		})
	}
	return arr
}

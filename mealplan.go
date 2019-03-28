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
var inputs []Input

func init() {
	flag.IntVar(&mealCount, "m", 10, "number of meals for the plan")
	// load the data and stuff
	inputFile, _ := os.Open("inputs.csv")
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
		inputs = append(inputs, Input{
			name:          line[0],
			category:      line[1],
			calories:      calories,
			quantity:      quantity,
			unit:          line[4],
			dryMultiplier: float32(dryMultiplier),
		})
	}
	// fmt.Println(inputs)
}

func main() {
	flag.Parse()
	generateMeals()
}

func generateMeals() {
	fmt.Println(fmt.Sprintf("gonna generate %d meals", mealCount))
	var meals []meal
	fmt.Println(meals)
	//
}

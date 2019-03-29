package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Input is an input in a recipe (basically an ingredient)
type Input struct {
	name          string
	category      string
	calories      float32
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
var styles []string

var mealLowerBound = 300
var mealUpperBound = 450

func init() {
	flag.IntVar(&mealCount, "m", 10, "number of meals for the plan")
	// load the data and stuff
	bases = loadData(bases, "data/bases.csv")
	fillings = loadData(fillings, "data/fillings.csv")
	styles = loadStyles(styles, "data/styles.txt")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	generateMeals()
}

// try is a var to prevent a stack overflow so that we only try 3 times before we give up
func generateMeal(base Input, filling Input, multiplier float32, try int) (string, float32) {
	if try == 0 {
		return "STOP", 1
	}
	var mealCalories = base.calories*multiplier + filling.calories
	log.Print(mealCalories)
	if float32(mealLowerBound) < mealCalories && mealCalories < float32(mealUpperBound) {
		log.Print("Found meal")
		return fmt.Sprintf(mealTitle, base.name, filling.name, styles[rand.Intn(len(styles))], 1), multiplier
	}
	if mealCalories < float32(mealUpperBound) {
		return generateMeal(base, filling, multiplier+0.25, try-1)
	}
	if float32(mealLowerBound) < mealCalories {
		log.Print("Decreasing base by 0.25")
		log.Print(float32(base.calories)*multiplier + float32(filling.calories))
		return generateMeal(base, filling, multiplier-0.25, try-1)
	}
	return "yo", multiplier
}
func generateMeals() {
	fmt.Println(fmt.Sprintf("gonna generate %d meals", mealCount))
	// var meals []meal
	var base = bases[rand.Intn(len(bases))]
	var filling = fillings[rand.Intn(len(fillings))]
	var meal, multiplier = generateMeal(base, filling, 1, 3)
	fmt.Println(meal)
	fmt.Printf(recipeLine, float32(base.quantity)*multiplier, base.unit, base.name, base.dryMultiplier*multiplier)
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
		calories, _ := strconv.ParseFloat(line[2], 32)
		quantity, _ := strconv.Atoi(line[3])
		dryMultiplier, _ := strconv.ParseFloat(line[5], 32)
		arr = append(arr, Input{
			name:          line[0],
			category:      line[1],
			calories:      float32(calories),
			quantity:      quantity,
			unit:          line[4],
			dryMultiplier: float32(dryMultiplier),
		})
	}
	return arr
}

func loadStyles(arr []string, fileName string) []string {
	inputFile, _ := os.Open(fileName)
	r := csv.NewReader(bufio.NewReader(inputFile))
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, line[0])
	}
	return arr
}

const mealTitle = `
1. %s and %s, %s style, %d servings
`

// 1 cups brown rice, or .33 cup dry
const recipeLine = `
%.2f %s %s, or %.2f cup dry
`
const calCount = `
%.2f calories per serving
`

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Input is an input in a recipe (basically an ingredient)
type Input struct {
	Name          string
	category      string
	calories      float32
	quantity      int
	Unit          string
	dryMultiplier float32
}

// Meal is a base and filling
type Meal struct {
	Base           Input
	Filling        Input
	Style          string
	Servings       int
	Calories       float32
	baseMultiplier float32
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
func generateMeal(base Input, filling Input, multiplier float32, try int) Meal {
	if try == 0 {
		return Meal{}
	}
	var mealCalories = base.calories*multiplier + filling.calories
	if float32(mealLowerBound) < mealCalories && mealCalories < float32(mealUpperBound) {
		return Meal{Base: base, Filling: filling, Style: styles[rand.Intn(len(styles))], baseMultiplier: multiplier}
	}
	if mealCalories < float32(mealUpperBound) {
		return generateMeal(base, filling, multiplier+0.25, try-1)
	}
	if float32(mealLowerBound) < mealCalories {
		return generateMeal(base, filling, multiplier-0.25, try-1)
	}
	return Meal{}
}
func generateMeals() {
	fmt.Println(fmt.Sprintf("%d meals you say? How about you make:", mealCount))
	var mealPlan []Meal
	var mealSplits []int
	mealSplits = splitMeals(mealSplits, mealCount)
	for _, servings := range mealSplits {
		var base = bases[rand.Intn(len(bases))]
		var filling = fillings[rand.Intn(len(fillings))]
		var meal = generateMeal(base, filling, 1, 3)
		meal.Servings = servings
		meal.Calories = meal.Base.calories*meal.baseMultiplier + meal.Filling.calories
		mealPlan = append(mealPlan, meal)
	}
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
	tmpl, err := template.New("mealplan").Funcs(funcMap).Parse(mealPlanTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, mealPlan)
	if err != nil {
		panic(err)
	}
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
			Name:          line[0],
			category:      line[1],
			calories:      float32(calories),
			quantity:      quantity,
			Unit:          line[4],
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

func splitMeals(arr []int, count int) []int {
	if count <= 3 {
		return append(arr, count)
	}
	if count%3 > 0 {
		return splitMeals(append(arr, 2), count-2)
	}
	return splitMeals(append(arr, 3), count-3)
}

const mealPlanTemplate = `{{range $i, $m := .}}
{{inc $i}}. {{$m.Base.Name}} and {{$m.Filling.Name}}, {{$m.Style}} style, {{$m.Servings}} servings{{end}}
{{range $i, $m := .}}
{{inc $i}}. {{$m.Base.Name}} and {{$m.Filling.Name}}, {{$m.Style}} style, {{$m.Servings}} servings
{{$m.Base.Unit}} {{$m.Base.Name}}, or X {{$m.Base.Unit}} dry
{{$m.Filling.Unit}} {{$m.Filling.Name}}, or X {{$m.Base.Unit}} dry
{{$m.Calories}} calories per serving
{{end}}
`

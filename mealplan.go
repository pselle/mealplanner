package main

import (
	"flag"
)

var mealCount int

func init() {
	flag.IntVar(&mealCount, "m", 10, "number of meals for the plan")
}

func main() {
	flag.Parse()
}

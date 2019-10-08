package main

import (
	"fmt"
	"os"
	"strconv"
)

// UNITS Maps between units
var UNITS = map[string]string{"celsius": "fahrenheit", "kilometers": "miles"}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use: converter <value> <unit>")
		os.Exit(1)
	}

	unitArg := os.Args[len(os.Args)-1]
	valueArg := os.Args[1 : len(os.Args)-1]

	unitTo, exists := UNITS[unitArg]
	if !exists {
		fmt.Printf("%s unit don't exist!", unitArg)
		os.Exit(1)
	}

	for i, v := range valueArg {
		valueArg, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("The value %s in the position %d not a valid number!\n", v, i)
			os.Exit(1)
		}

		var valueTo float64

		if unitArg == "celsius" {
			valueTo = valueArg*1.8 + 32
		} else {
			valueTo = valueArg / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			valueArg, unitArg, valueTo, unitTo)
	}
}

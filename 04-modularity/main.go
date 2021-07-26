package main

import (
	"fmt"
	calc "modularity/calculator"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	color.Red(strconv.Itoa(calc.Add(10, 20)))
	color.Green(strconv.Itoa(calc.Sub(10, 20)))
	color.Yellow(fmt.Sprintf("Invocation count = %d\n", calc.GetInvocationCount()))
	color.Red(strconv.Itoa(calc.Add(100, 20)))
	color.Yellow(fmt.Sprintf("Invocation count = %d\n", calc.GetInvocationCount()))
}

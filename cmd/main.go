package main

import (
	"fmt"

	"github.com/nerfthisdev-itmo/cm-lab-5/internal/interpol"
)

func main() {

	testvalues := interpol.FuncValues{
		X: []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		Y: []float64{1.25, 2.38, 3.79, 5.44, 7.14},
	}

	table := interpol.BuildFiniteDifferenceTable(testvalues)

	interpol.PrintFiniteDifferenceTable(table, testvalues)

	fmt.Printf("Интерполяция многочленом Лагранжа: %f\n", interpol.Lagrange(testvalues, 0.35))
}

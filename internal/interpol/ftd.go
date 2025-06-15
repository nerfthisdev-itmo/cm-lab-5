package interpol

import "fmt"

type FiniteDifferenceTable struct {
	Table [][]float64
}

func BuildFiniteDifferenceTable(values FuncValues) FiniteDifferenceTable {
	n := len(values.Y)
	table := make([][]float64, n)

	for i := range n {
		table[i] = make([]float64, n-i)
		table[i][0] = values.Y[i]
	}

	for j := 1; j < n; j++ {
		for i := range n - j {
			table[i][j] = table[i+1][j-1] - table[i][j-1]
		}
	}

	return FiniteDifferenceTable{Table: table}
}

func BuildDividedDifferenceTable(values FuncValues) [][]float64 {
	n := len(values.X)
	table := make([][]float64, n)

	for i := range n {
		table[i] = make([]float64, n)
		table[i][0] = values.Y[i]
	}

	for j := 1; j < n; j++ {
		for i := range n - j {
			table[i][j] = (table[i+1][j-1] - table[i][j-1]) / (values.X[i+j] - values.X[i])
		}
	}

	return table
}

func PrintFiniteDifferenceTable(table FiniteDifferenceTable, values FuncValues) {
	fmt.Printf("%-10s", "x")
	for i := range table.Table[0] {
		fmt.Printf("%-12s", fmt.Sprintf("Δ^%d y", i))
	}
	fmt.Println()

	for i := range table.Table {
		fmt.Printf("%-10.4f", values.X[i])
		for j := range table.Table[i] {
			fmt.Printf("%-12.4f", table.Table[i][j])
		}
		fmt.Println()
	}
}

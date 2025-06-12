package interpol

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
		for i := range n-j { 
			table[i][j] = table[i+1][j-1] - table[i][j-1]
		}
	}

	return FiniteDifferenceTable{Table: table}
}

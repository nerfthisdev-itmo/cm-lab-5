package interpol

func NewtonForwardDifference(table FiniteDifferenceTable, values FuncValues, x float64) float64 {
	n := len(values.X)
	h := values.X[1] - values.X[0]
	t := (x - values.X[0]) / h

	result := table.Table[0][0]
	prod := 1.0

	for i := 1; i < n; i++ {
		prod *= (t - float64(i-1))
		result += (prod / factorial(i)) * table.Table[0][i]
	}

	return result
}

func NewtonDividedDifference(table DividedDifferenceTable, values FuncValues, x float64) float64 {
	n := len(values.X)
	result := table.Table[0][0]
	prod := float64(1)

	for i := 1; i < n; i++ {
		prod *= (x - values.X[i-1])
		result += table.Table[0][i] * prod
	}
	return result
}

func factorial(n int) float64 {
	res := 1.0
	for i := 2; i <= n; i++ {
		res *= float64(i)
	}
	return res
}

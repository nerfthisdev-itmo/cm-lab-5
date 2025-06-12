package interpol

func Lagrange(values FuncValues, x float64) float64 {

	n := len(values.X)
	result := 0.0

	for i := range n {
		term := values.Y[i]
		for j := range n {
			if j != i {
				term *= (x - values.X[j]) / (values.X[i] - values.X[j])
			}
		}
		result += term
	}

	return result

}

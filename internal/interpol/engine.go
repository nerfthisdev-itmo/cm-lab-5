package interpol

import "fmt"

func ApplyAllInterpolations(values FuncValues, x float64) {
	fmt.Printf("\nИнтерполяция в точке x = %.4f\n", x)
	fdt := BuildFiniteDifferenceTable(values)
	yFinite := NewtonForwardDifference(fdt, values, x)
	fmt.Printf("Многочлен Ньютона (Конечные разности): %.6f\n", yFinite)

	ddt := BuildDividedDifferenceTable(values)
	yDivided := NewtonDividedDifference(ddt, values, x)
	fmt.Printf("Многочлен Ньютона (разделенные разности): %.6f\n", yDivided)

	lagrange := Lagrange(values, x)
	fmt.Printf("Многочлен Лагранжа: %.6f\n", lagrange)
}

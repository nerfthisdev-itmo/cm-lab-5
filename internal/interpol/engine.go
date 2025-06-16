package interpol

import (
	"fmt"
	"math"
)

func ApplyAllInterpolations(values FuncValues, x float64) {
	fmt.Printf("\nИнтерполяция в точке x = %.4f\n", x)

	uniform := true
	if len(values.X) > 1 {
		h := values.X[1] - values.X[0]
		for i := 2; i < len(values.X); i++ {
			if math.Abs(values.X[i]-values.X[i-1]-h) > 1e-9 {
				uniform = false
				break
			}
		}
	}

	if uniform {
		fmt.Println("Шаг h постоянный, используем метод конечных разностей.")
		fdt := BuildFiniteDifferenceTable(values)
		PrintFiniteDifferenceTable(fdt, values)
		yFinite := NewtonForwardDifference(fdt, values, x)
		fmt.Printf("Многочлен Ньютона (конечные разности): %.6f\n", yFinite)
	} else {
		fmt.Println("Шаг h непостоянный, метод по конечным разностям не применяется.")
		ddt := BuildDividedDifferenceTable(values)
		PrintDividedDifferenceTable(ddt, values)
		yDivided := NewtonDividedDifference(ddt, values, x)
		fmt.Printf("Многочлен Ньютона (разделённые разности): %.6f\n", yDivided)
	}

	lagrange := Lagrange(values, x)
	fmt.Printf("Многочлен Лагранжа: %.6f\n", lagrange)
}

package main

import (
	"fmt"

	"github.com/nerfthisdev-itmo/cm-lab-5/internal/interpol"
	"github.com/nerfthisdev-itmo/cm-lab-5/internal/utils"
)

func main() {
	values, err := utils.HandleInput()
  

	if err != nil {
		panic(err)
	}

	fmt.Println(values.X, values.Y)

  table := interpol.BuildFiniteDifferenceTable(values)

  interpol.PrintFiniteDifferenceTable(table, values)
}

package main

import (
	"fmt"

	"github.com/nerfthisdev-itmo/cm-lab-5/internal/utils"
)

func main() {
	values, err := utils.HandleInput()
  

	if err != nil {
		panic(err)
	}

	fmt.Println(values.X, values.Y)
}

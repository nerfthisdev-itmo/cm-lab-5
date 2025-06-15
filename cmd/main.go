package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nerfthisdev-itmo/cm-lab-5/internal/interpol"
	"github.com/nerfthisdev-itmo/cm-lab-5/internal/utils"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	values, err := utils.HandleInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Введите значение x, в котором нужно интерполировать:")
	fmt.Print("> ")
	xLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка чтения x: ", err)
	}
	x, err := strconv.ParseFloat(strings.TrimSpace(xLine), 64)
	if err != nil {
		log.Fatal("Ошибка преобразования x: ", err)
	}

	interpol.ApplyAllInterpolations(values, x)
}

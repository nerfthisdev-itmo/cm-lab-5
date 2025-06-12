package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nerfthisdev-itmo/cm-lab-5/internal/interpol"
)

var reader = bufio.NewReader(os.Stdin)

func HandleInput() (interpol.FuncValues, error) {
	fmt.Println("Как вы хотите осуществить ввод?")
	fmt.Println("Ручной / Файл / Выбор функции (i/f/c)")
	fmt.Print("> ")

	choiceLine, err := reader.ReadString('\n')
	if err != nil {
		return interpol.FuncValues{}, fmt.Errorf("ошибка чтения ввода: %w", err)
	}
	choice := strings.ToLower(strings.TrimSpace(choiceLine))

	if len(choice) == 0 {
		return interpol.FuncValues{}, errors.New("не введён режим ввода")
	}

	switch choice[0] {
	case 'i':
		return manualInput()
	case 'f':
		return fileInput()
	case 'c':
		return interpol.FuncValues{}, chooseFunction()
	default:
		return interpol.FuncValues{}, fmt.Errorf("неизвестный тип ввода: %s", choice)
	}
}

func manualInput() (interpol.FuncValues, error) {

	fmt.Println("Введите значения x (через пробел)")
	fmt.Print("> ")
	xLine, err := reader.ReadString('\n')
	if err != nil {
		return interpol.FuncValues{}, fmt.Errorf("ошибка чтения x: %w", err)
	}

	fmt.Println("Введите значения y (через пробел)")
	fmt.Print("> ")
	yLine, err := reader.ReadString('\n')
	if err != nil {
		return interpol.FuncValues{}, fmt.Errorf("ошибка чтения y: %w", err)
	}

	xs := strings.Fields(xLine)
	ys := strings.Fields(yLine)

	if len(xs) != len(ys) {
		return interpol.FuncValues{}, fmt.Errorf("количество значений x (%d) не совпадает с количеством y (%d)", len(xs), len(ys))
	}

	fmt.Println("Ввод принят.")

	xsFloat, err := parseFloatSlice(xs)

	if err != nil {
		return interpol.FuncValues{}, err
	}

	ysFloat, err := parseFloatSlice(ys)

	if err != nil {
		return interpol.FuncValues{}, err
	}

	return interpol.FuncValues{X: xsFloat, Y: ysFloat}, nil

}

func fileInput() (interpol.FuncValues, error) {
	fmt.Println("Введите путь к файлу")
	fmt.Print("> ")
	filePathLine, err := reader.ReadString('\n')
	if err != nil {
		return interpol.FuncValues{}, fmt.Errorf("ошибка чтения пути: %w", err)
	}
	filePath := strings.TrimSpace(filePathLine)

	if filePath == "" {
		return interpol.FuncValues{}, errors.New("путь к файлу не может быть пустым")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return interpol.FuncValues{}, err
	}
	defer file.Close()

	fmt.Println("файл успешно открыт. Содержимое:")

	scanner := bufio.NewScanner(file)
	var line1, line2 string

	if scanner.Scan() {
		line1 = scanner.Text()
	} else {
		fmt.Println("Файл пуст или первая строка отсутствует")
		return interpol.FuncValues{}, fmt.Errorf("Файл пуст или первая строка отсутствует")
	}

	if scanner.Scan() {
		line2 = scanner.Text()
	} else {
		return interpol.FuncValues{}, fmt.Errorf("В файле только одна строка")

	}

	xs := strings.Fields(line1)
	ys := strings.Fields(line2)
	xsFloat, err := parseFloatSlice(xs)

	if err != nil {
		return interpol.FuncValues{}, err
	}

	ysFloat, err := parseFloatSlice(ys)

	if err != nil {
		return interpol.FuncValues{}, err
	}

	return interpol.FuncValues{X: xsFloat, Y: ysFloat}, nil
}

func chooseFunction() error {
	fmt.Println("Вы выбрали выбор функции. (не реализовано)")
	return nil
}

func parseFloatSlice(s []string) ([]float64, error) {
	si := make([]float64, 0, len(s))
	for _, a := range s {
		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return si, err
		}
		si = append(si, i)

	}
	return si, nil
}

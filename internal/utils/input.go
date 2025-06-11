package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func HandleInput() error {
	fmt.Println("Как вы хотите осуществить ввод?")
	fmt.Println("Ручной / Файл / Выбор функции (i/f/c)")
	fmt.Print("> ")

	choiceLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("ошибка чтения ввода: %w", err)
	}
	choice := strings.ToLower(strings.TrimSpace(choiceLine))

	if len(choice) == 0 {
		return errors.New("не введён режим ввода")
	}

	switch choice[0] {
	case 'i':
		return manualInput()
	case 'f':
		return fileInput()
	case 'c':
		return chooseFunction()
	default:
		return fmt.Errorf("неизвестный тип ввода: %s", choice)
	}
}

func manualInput() error {
	fmt.Println("Введите значения x (через пробел)")
	fmt.Print("> ")
	xLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("ошибка чтения x: %w", err)
	}

	fmt.Println("Введите значения y (через пробел)")
	fmt.Print("> ")
	yLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("ошибка чтения y: %w", err)
	}

	xs := strings.Fields(xLine)
	ys := strings.Fields(yLine)

	if len(xs) != len(ys) {
		return fmt.Errorf("количество значений x (%d) не совпадает с количеством y (%d)", len(xs), len(ys))
	}

	fmt.Println("Ввод принят.")
	return nil
}

func fileInput() error {
	fmt.Println("Введите путь к файлу")
	fmt.Print("> ")
	filePathLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("ошибка чтения пути: %w", err)
	}
	filePath := strings.TrimSpace(filePathLine)

	if filePath == "" {
		return errors.New("путь к файлу не может быть пустым")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	fmt.Println("файл успешно открыт. Содержимое:")

	scanner := bufio.NewScanner(file)
	var line1, line2 string

	if scanner.Scan() {
		line1 = scanner.Text()
	} else {
		fmt.Println("Файл пуст или первая строка отсутствует")
		return fmt.Errorf("Файл пуст или первая строка отсутствует")
	}

	if scanner.Scan() {
		line2 = scanner.Text()
	} else {
		return fmt.Errorf("В файле только одна строка")

	}

	xs := strings.Fields(line1)
	ys := strings.Fields(line2)
	fmt.Printf("x: %v \n", xs)
	fmt.Printf("y: %v \n", ys)
	return nil
}

func chooseFunction() error {
	fmt.Println("Вы выбрали выбор функции. (не реализовано)")
	return nil
}

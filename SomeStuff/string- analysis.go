/*
Задача — разработать программу, которая будет анализировать текст, введенный пользователем.
Программа должна выполнять несколько ключевых функций, позволяющих получить различные статистические данные о тексте.

Ввод текста:
Создайте функцию GetInput() (string, error), которая запрашивает у пользователя ввод текста и возвращает его для дальнейшего анализа. Если функция вернет ошибку, программу необходимо завершить.

Подсчет символов:
Реализуйте функцию CountCharacters(text string) (letters int, digits int, spaces int, punctuation int), которая принимает текст и возвращает количество:
- букв (letters)
- цифр (digits)
- пробелов (spaces)
- знаков препинания (punctuation)

Вывод результатов:
Напишите функцию DisplayResults(letters, digits, spaces, punctuation int), которая принимает результаты анализа и выводит их на экран в удобочитаемом формате. 

Например:
Количество букв: 50
Количество цифр: 10
Количество пробелов: 5
Количество знаков препинания: 3
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func CountCharacters(text string) (letters int, digits int, spaces int, punctuation int) {
	for _, char := range text {
		//fmt.Println(string(char))
		switch {
		case char >= '0' && char <= '9':
			digits++
		case char == ' ':
			spaces++
		case unicode.IsPunct(char):
			punctuation++
		case unicode.IsLetter(char):
			letters++
		}
	}
	return
}

func GetInput() (string, error) {
	fmt.Printf("Программа анализирует введенный пользователем текст и выводит следующую информацию:\n")
	fmt.Printf("- количество букв\n- количество цифр\n- количество пробелов\n- количество знаков препинания\n\n")
	fmt.Printf("Введите текст: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println()
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("ошибка сканирования: %w", err)
		}
		return "", errors.New("невозможно прочитать инпут")
	}
	return scanner.Text(), nil
}

func DisplayResults(letters int, digits int, spaces int, punctuation int) {
	fmt.Printf("Количество букв: %d\n", letters)
	fmt.Printf("Количество цифр: %d\n", digits)
	fmt.Printf("Количество пробелов: %d\n", spaces)
	fmt.Printf("Количество знаков препинания: %d\n", punctuation)
}

func main() {
	input, err := GetInput()
	if err != nil {
		fmt.Printf("Ошибка при ввводе текста: %v\n", err)
		os.Exit(1)
	} else {
		letters, digits, spaces, punctuation := CountCharacters(input)
		DisplayResults(letters, digits, spaces, punctuation)
	}
}

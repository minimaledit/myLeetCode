package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	keyboard := "Клавиатура JZ9: 19200"
	headphones := "Наушники N45: 9600"
	smartphone := "Смартфон S10: 55000"

	//Input
	fmt.Print("Введите название товара: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Oшибка ввода: %s", err)
	}
	userInput := scanner.Text()

  
	userInput = strings.ToLower(userInput)

	if strings.Contains(strings.ToLower(keyboard), userInput) {
		fmt.Println(keyboard)
	} else if strings.Contains(strings.ToLower(headphones), userInput) {
		fmt.Println(headphones)
	} else if strings.Contains(strings.ToLower(smartphone), userInput) {
		fmt.Println(smartphone)
	} else {
		fmt.Printf("Товар \"%s\" не найден.\n", userInput)
	}
}

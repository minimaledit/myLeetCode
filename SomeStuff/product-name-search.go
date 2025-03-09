/*
Напишите программу, которая будет выполнять поиск товара по его названию. 
Пользователь должен ввести название товара или его часть, после чего программа должна вывести в консоль цену данного товара в формате:
Название товара: цена товара

Доступные товары и их цены:
"Клавиатура JZ9": 19200
"Наушники N45": 9600
"Смартфон S10": 55000

Если пользователь введет название товара, которого нет в списке, программа должна вывести сообщение:
Товар "название_введенного_товара" не найден.

Поиск должен быть нечувствителен к регистру. Например, если пользователь введет "наУШНИКИ n45", программа должна вернуть цену 9600.
*/

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

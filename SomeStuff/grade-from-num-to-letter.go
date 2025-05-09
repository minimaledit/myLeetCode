/*
Нашему проекту приходится работать с разными учебными заведениями и не всегда система оценок похожа.
Необходимо создать программу, которая запрашивает у пользователя оценку в диапазоне от 0 до 100 включительно и выводит соответствующую буквенную оценку (A, B, C, D, F).

Необходимо создать функцию, которая будет запрашивать у пользователя ввод числовой оценки (целое число) в диапазоне от 0 до 100.
Необходимо убедиться, что введенное значение находится в указанном диапазоне.
Если значение вне диапазона или введено некорректно, программа должна возвращать ошибку из функции, а при ее обработке выводить сообщение об ошибке и завершать выполнение.

В зависимости от введенной оценки, программа должна присвоить ей соответствующую буквенную оценку по следующей шкале:
- 90–100: A
- 80–89: B
- 70–79: C
- 60–69: D
- Ниже 60: F
*/

import (
	"errors"
	"fmt"
)

func score() (int, error) {
	var score int
	fmt.Printf("Введите вашу отметку [0-100]: ")
	_, err := fmt.Scan(&score)
	if err != nil {
		return 0, errors.New("ошибка ввода: введите целое число")
	}

	if score < 0 || score > 100 {
		return 0, errors.New("ошибка ввода: введите число из диапазона от 0 до 100")
	}

	return score, err
}

func letterScore(numMark int) (numWord string) {
	switch {
	case numMark < 60:
		return "F"
	case numMark < 70:
		return "D"
	case numMark < 80:
		return "C"
	case numMark < 90:
		return "B"
	default:
		return "A"
	}
}

func main() {
	score, err := score()
	if err != nil {
		fmt.Println(err)
		return
	}
	letterGrade := letterScore(score)
	fmt.Printf("Ваша буквенная оценка: %s\n", letterGrade)
}

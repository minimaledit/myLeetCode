/*
Декорируем ролевую игру

Необходимо создать некоторые функции для текстовой ролевой игры, в которой игроки могут использовать различные способности.
Ваша задача заключается в реализации декораторов, которые добавляют уникальные эффекты к базовым способностям персонажа, используя только функции.

Базовая способность:
Реализуйте функцию Attack() string, которая будет представлять базовую способность персонажа и возвращать строку "Атака выполнена!".

Декоратор увеличения урона:
Создайте функцию DamageBoostDecorator(attackFunc func() string) func() string, которая принимает функцию атаки и возвращает новую функцию.
Эта новая функция должна выводить сообщение "Вам улыбнулась удача, нанесение урона увеличено на 10%!" перед вызовом базовой функции атаки.

Декоратор критического удара:
Создайте функцию CriticalHitDecorator(attackFunc func() string) func() string, которая добавляет шанс критического удара.
Если критический удар происходит (например, с вероятностью 25%), возвращайте строку "Критический удар! Урон удвоен!". В противном случае вызывайте функцию атаки.

Декоратор эффекта замедления:
Создайте функцию SlowEffectDecorator(attackFunc func() string) func() string, которая добавляет эффект замедления к атаке.
Например, возвращайте строку "Цель замедлена на 2 хода!" после выполнения функции атаки.

Комбинирование декораторов:
Позвольте игроку комбинировать декораторы. Например, игрок может использовать атаку с увеличением урона и критическим ударом одновременно.
*/
import (
	"fmt"
	"math/rand/v2"
)

// Базовая способность
func Attack() string {
	return "Атака выполнена!"
}

// Декоратор увеличения урона
func DamageBoostDecorator(attackFunc func() string) func() string {
	return func() string {
		return "Вам улыбнулась удача, нанесение урона увеличено на 10%! " + attackFunc()
	}
}

// Декоратор критического удара
func CriticalHitDecorator(attackFunc func() string) func() string {
	return func() string {
		chance := rand.IntN(100)
		if chance < 25 {
			return "Критический удар! Урон удвоен! " + attackFunc()
		}
		return attackFunc()
	}
}

// Декоратор эффекта замедления
func SlowEffectDecorator(attackFunc func() string) func() string {
	return func() string {
		return attackFunc() + " Цель замедлена на 2 хода!"
	}
}

// Базовая атака
func basicAttack() string {
	return "Вы нанесли 100 урона!"
}

func main() {
	fmt.Println(Attack())

	attackWithDamageBoost := DamageBoostDecorator(Attack)
	fmt.Println(attackWithDamageBoost())

	attackWithCriticalHit := CriticalHitDecorator(Attack)
	fmt.Println(attackWithCriticalHit())

	attackWithSlowEffect := SlowEffectDecorator(Attack)
	fmt.Println(attackWithSlowEffect())

	combinedAttack := SlowEffectDecorator(CriticalHitDecorator(DamageBoostDecorator(Attack)))
	fmt.Println(combinedAttack())
}

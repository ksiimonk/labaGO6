package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для расчёта факториала числа n
func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Факториал %d = %d\n", n, result)
}

// Функция для генерации случайных чисел
func generateRandomNumbers(count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Случайное число: %d\n", num)
		time.Sleep(150 * time.Millisecond) // Имитируем задержку
	}
}

// Функция для вычисления суммы ряда чисел от 1 до n
func sumSeries(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(200 * time.Millisecond) // Имитируем задержку
	}
	fmt.Printf("Сумма ряда от 1 до %d = %d\n", n, sum)
}

func main() {
	// Запуск трёх функций в горутинах
	go factorial(5)
	go generateRandomNumbers(5)
	go sumSeries(5)

	// Даем время горутинам завершить выполнение
	time.Sleep(3 * time.Second)
	fmt.Println("Все горутины завершены.")
}

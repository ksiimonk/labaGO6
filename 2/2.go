package main

import (
	"fmt"
	"time"
)

// Функция для генерации первых n чисел Фибоначчи и отправки их в канал
func generateFibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(100 * time.Millisecond) // Имитируем задержку
	}
	close(ch) // Закрываем канал, когда все данные отправлены
}

// Функция для чтения данных из канала и вывода их на экран
func printFibonacci(ch chan int) {
	// Используем цикл для чтения данных из канала
	// Канал автоматически завершает цикл, когда он закрыт
	for num := range ch {
		fmt.Println(num)
		time.Sleep(150 * time.Millisecond) // Имитируем задержку
	}
}

func main() {
	// Создаём канал для передачи целых чисел
	fibChannel := make(chan int)

	// Запускаем горутины для генерации и вывода чисел
	go generateFibonacci(10, fibChannel)
	go printFibonacci(fibChannel)

	// Даем время горутинам завершить выполнение
	time.Sleep(2 * time.Second)
	fmt.Println("Все горутины завершены.")
}

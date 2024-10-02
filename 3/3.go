package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Горутина для генерации случайных чисел и отправки их в канал
func generateRandomNumbers(numbers chan int) {
	for {
		num := rand.Intn(100) // Генерируем случайное число от 0 до 99
		numbers <- num
		time.Sleep(200 * time.Millisecond) // Задержка для имитации работы
	}
}

// Горутина для проверки чётности/нечётности числа и отправки сообщения в другой канал
func checkEvenOdd(numbers chan int, messages chan string) {
	for {
		num := <-numbers // Получаем число из канала
		if num%2 == 0 {
			messages <- fmt.Sprintf("Число %d — чётное", num)
		} else {
			messages <- fmt.Sprintf("Число %d — нечётное", num)
		}
		time.Sleep(100 * time.Millisecond) // Задержка для имитации работы
	}
}

func main() {
	// Создаём два канала
	numbers := make(chan int)
	messages := make(chan string)

	// Запускаем горутины
	go generateRandomNumbers(numbers)
	go checkEvenOdd(numbers, messages)

	// Используем select для одновременной обработки данных из двух каналов
	for {
		select {
		case num := <-numbers:
			fmt.Printf("Сгенерировано число: %d\n", num)
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}

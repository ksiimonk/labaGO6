package main

import (
	"fmt"
	"sync"
)

// Общая переменная-счётчик
var counter = 0

// Мьютекс для синхронизации доступа к счётчику
var mutex sync.Mutex

// Функция, которая увеличивает счётчик
func increment(wg *sync.WaitGroup, useMutex bool) {
	defer wg.Done() // Сообщаем WaitGroup, что горутина завершена

	for i := 0; i < 1000; i++ {
		if useMutex {
			// Блокируем мьютекс перед изменением счётчика
			mutex.Lock()
		}

		// Увеличиваем счётчик
		counter++

		if useMutex {
			// Разблокируем мьютекс после изменения счётчика
			mutex.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Запуск горутин с использованием мьютекса
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&wg, true) // Включён мьютекс
	}
	wg.Wait()

	fmt.Printf("Счётчик с мьютексом: %d\n", counter)

	// Обнуляем счётчик
	counter = 0

	// Запуск горутин без использования мьютекса
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&wg, false) // Мьютекс выключен
	}
	wg.Wait()

	fmt.Printf("Счётчик без мьютекса: %d\n", counter)
}

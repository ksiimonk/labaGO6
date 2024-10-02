package main

import (
	"fmt"
	"log"
	"sync"
)

// Структура для запроса на выполнение операции
type CalculationRequest struct {
	A, B   float64      // Операнды
	Op     string       // Операция: "+", "-", "*", "/"
	Result chan float64 // Канал для возврата результата
	Error  chan error   // Канал для возврата ошибки (например, деление на 0)
}

// Серверная часть калькулятора, которая выполняет операции
func calculator(requests chan CalculationRequest, wg *sync.WaitGroup) {
	defer wg.Done() // Уведомляем о завершении горутины

	for req := range requests {
		var res float64
		var err error

		// Выполнение операции на основе типа
		switch req.Op {
		case "+":
			res = req.A + req.B
		case "-":
			res = req.A - req.B
		case "*":
			res = req.A * req.B
		case "/":
			if req.B == 0 {
				err = fmt.Errorf("деление на ноль")
			} else {
				res = req.A / req.B
			}
		default:
			err = fmt.Errorf("неизвестная операция: %s", req.Op)
		}

		// Отправляем результат или ошибку в соответствующий канал
		if err != nil {
			req.Error <- err
		} else {
			req.Result <- res
		}
	}
}

func main() {
	// Создаём канал для отправки запросов
	requests := make(chan CalculationRequest)

	// WaitGroup для синхронизации завершения работы калькулятора и клиентских горутин
	var wg sync.WaitGroup

	// Запускаем серверную часть калькулятора в отдельной горутине
	wg.Add(1)
	go calculator(requests, &wg)

	// Создаём клиентские запросы
	operations := []CalculationRequest{
		{A: 10, B: 20, Op: "+", Result: make(chan float64), Error: make(chan error)},
		{A: 15, B: 5, Op: "-", Result: make(chan float64), Error: make(chan error)},
		{A: 7, B: 3, Op: "*", Result: make(chan float64), Error: make(chan error)},
		{A: 40, B: 0, Op: "/", Result: make(chan float64), Error: make(chan error)}, // Ошибка деления на 0
		{A: 16, B: 4, Op: "/", Result: make(chan float64), Error: make(chan error)},
	}

	// Используем WaitGroup для клиентских горутин
	var clientWg sync.WaitGroup

	// Запуск клиентских запросов в отдельных горутинах
	for _, op := range operations {
		clientWg.Add(1)
		go func(op CalculationRequest) {
			defer clientWg.Done()

			// Отправляем запрос на выполнение
			requests <- op

			// Ожидаем либо результат, либо ошибку
			select {
			case res := <-op.Result:
				fmt.Printf("Результат %f %s %f = %f\n", op.A, op.Op, op.B, res)
			case err := <-op.Error:
				log.Printf("Ошибка при выполнении операции %f %s %f: %s\n", op.A, op.Op, op.B, err)
			}
		}(op)
	}

	// Ожидаем завершения всех клиентских горутин
	clientWg.Wait()

	// Закрываем канал с запросами после того, как все клиентские горутины завершились
	close(requests)

	// Ожидаем завершения работы калькулятора
	wg.Wait()

	fmt.Println("Все операции завершены.")
}

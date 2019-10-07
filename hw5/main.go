package main

import (
	"fmt"
	"os"
	"time"
)

func work(tasks []func()) {
	// Стартуем подсчет ошибок
	go func(err *int) {
		for x := range Errors {
			*err = *err + x
			fmt.Printf("error count: %d\n", ErrorCount)
		}
	}(&ErrorCount)

	for _, f := range tasks {
		// Записываем в канал с Max-размером N, если мы его заполним,
		// то он заблокируется на запись до завершения одной из горутин
		WorkingTasks <- 1

		if ErrorCount == MaxErrorCount {
			fmt.Println("Max error count occurred")
			return
		}

		go f()
	}

}

// MaxWorkingGoroutines - максимальное кол-во работающих горутин
var MaxWorkingGoroutines = 1

// MaxErrorCount - максимальное кол-во ошибок
var MaxErrorCount = 1

// ErrorCount - кол-во ошибок
var ErrorCount = 0

// WorkingTasks - канал, который хранит текущее кол-во запущенных горутин
var WorkingTasks = make(chan int, MaxWorkingGoroutines)

// Errors - канал, через который приходит инфа об ошибке
var Errors = make(chan int)

func main() {

	tasks := []func(){
		func() {
			fmt.Println("f1")
			Errors <- 1
			time.Sleep(1 * time.Second)

			x := <-WorkingTasks
			fmt.Printf("wt: %d\n", x)
		},
		func() {
			fmt.Println("f2")
			time.Sleep(1 * time.Second)
			x := <-WorkingTasks
			fmt.Printf("wt: %d\n", x)
		}}

	work(tasks)

	time.Sleep(3 * time.Second)

	close(WorkingTasks)
	close(Errors)

	os.Exit(0)
}

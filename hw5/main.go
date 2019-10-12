package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func work(tasks []func() error) {
	for i := 0; i < MaxWorkingGoroutines; i++ {
		go func() {
			for {
				f, _ := <-functions
				if len(errored) == MaxErrorCount {
					fmt.Println("max errors exceeded")
					return
				}

				if err := f(); err != nil {
					fmt.Println("error detected")
					errored <- 1
				}
			}
		}()
	}

	for _, f := range tasks {
		functions <- f
	}
}

// MaxWorkingGoroutines - максимальное кол-во работающих горутин
var MaxWorkingGoroutines = 2

// MaxErrorCount - максимальное кол-во ошибок
var MaxErrorCount = 1

// ErrorCount - кол-во ошибок
//var ErrorCount = 0

var functions = make(chan func() error, MaxWorkingGoroutines)

var errored = make(chan int, MaxErrorCount)

func main() {

	tasks := []func() error{
		func() error {
			fmt.Println("f1")
			time.Sleep(1 * time.Second)
			return errors.New("error1")
		},
		func() error {
			fmt.Println("f2")
			time.Sleep(1 * time.Second)
			return nil
		},
		func() error {
			fmt.Println("f3")
			time.Sleep(1 * time.Second)
			return errors.New("error2")
		},
		func() error {
			fmt.Println("f4")
			time.Sleep(1 * time.Second)
			return nil
		}}

	work(tasks)

	for {
		if len(functions) == 0 {
			break
		}

		if len(errored) == MaxErrorCount {
			break
		}

		time.Sleep(time.Second)
		fmt.Println("wait for finish")
	}

	close(errored)
	close(functions)

	os.Exit(0)
}

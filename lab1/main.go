package main

import (
	"errors"
	"fmt"
)

func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница диапазона больше правой")
	}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func main() {
	fmt.Println(hello("Андрей"))

	err := printEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = printEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

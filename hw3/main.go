package main

import (
	"fmt"

	"hw3/stack"
)

func main() {
	s := stack.New[int]()

	fmt.Println("Пуст ли стек:", s.IsEmpty())
	fmt.Println("Текущий размер стека:", s.Size())

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("После добавления элементов:")
	fmt.Println("Размер стека:", s.Size())
	fmt.Println("Пуст ли стек:", s.IsEmpty())

	if v, ok := s.Pop(); ok {
		fmt.Println("Извлечён элемент:", v)
	}
	if v, ok := s.Pop(); ok {
		fmt.Println("Извлечён элемент:", v)
	}
	if v, ok := s.Pop(); ok {
		fmt.Println("Извлечён элемент:", v)
	}

	fmt.Println("После извлечения всех элементов:")
	fmt.Println("Размер стека:", s.Size())
	fmt.Println("Пуст ли стек:", s.IsEmpty())

	s.Push(42)
	s.Push(99)
	fmt.Println("Перед очисткой:")
	fmt.Println("Размер стека:", s.Size())

	s.Clear()
	fmt.Println("После очистки:")
	fmt.Println("Размер стека:", s.Size())
	fmt.Println("Пуст ли стек:", s.IsEmpty())
}

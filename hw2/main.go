package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	if num >= 12307 {
		fmt.Println("Введите число меньше 12307")
		return
	} else {
		num = Oper(num)
	}
	fmt.Printf("Супер, после небольших махинаций мы получили число %d", num)
}

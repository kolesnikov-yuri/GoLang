package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		age  int
		name string
	)

	// fmt.Scan(&age)  15
	// fmt.Scan(&name)  Bob
	fmt.Scan(&age, &name)

	fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)
	// My name is: Bob
	// My age is: 15

	//Для ручного использования потока ввода
	fmt.Fscan(os.Stdin, &age)    // 16
	fmt.Println("New age:", age) // New age: 16

}

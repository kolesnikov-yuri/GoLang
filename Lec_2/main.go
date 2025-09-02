package main //package name

import (
	"fmt"
	"math"
)

func main() {

	//Простейший вывод на косноль. println - это вывод аргумента + '\n'
	fmt.Println("Это моя вторая программа! Я рад, что учусь здесь!!")
	// Это моя вторая программа! Я рад, что учусь здесь!!
	fmt.Println("Second line")
	// Second line

	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	// FirstSecondThird

	// 	//Форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
	// Hello, my name is Bob
	// My age is 42

	/////////////////////////////////
	///////////////////////////////////

	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age) // My age is: 0
	age = 32
	fmt.Println("Age after assignment:", age) // Age after assignment: 32

	// 	//Декларирование и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height) // My height is: 183

	// 	//В чем "полустрогость" типзации? Можно опускать тип переменной
	var uid = 12345
	fmt.Println("My uid:", uid) // My uid: 12345
	fmt.Printf("%T\n", uid)     // int показывает тип переменной

	// 	//Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar) // FirstVar:20 SecondVar:30

	//Декларирвоание блока переменных
	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge %d\nUID: %d\n", personName, personAge, personUID)
	// Name: Bob
	// Age 42
	// UID: 0

	//Немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b) // 30 Vova

	a = 300
	//Немного хорошего. Повторное декларирование переменной приводит к ошибке компиляции
	// var a = 200
	// fmt.Print(a) // ошибка .\main.go:67:6: a redeclared in this block
	// .\main.go:62:6: other declaration of a

	//Короткая декларция (корткое объявление)
	count := 10
	fmt.Println("Count:", count) // Count: 10
	count = count + 1
	fmt.Println("Count:", count) // Count: 11

	//Множественное присваивание через :=
	// aArg, bArg := 10, "Vova"
	// fmt.Println(aArg, bArg) // 10 Vova

	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg) // 10 30
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg) // 30 40

	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg) // ошибка (передиклорация)

	//Исключение из этого правила
	bArg, cArg := 300, 400
	fmt.Println(aArg, bArg, cArg) // 30 300 400 ошибки нет + новая переменная

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.2f\n", math.Min(width, length))
	// Min dimensional of rectangle is : 20.50
}

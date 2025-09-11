package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean) // fals (по умолчанию fals)

	//Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean) // fals
	fmt.Println("OR:", aBoolean || bBoolean)  // true
	fmt.Println("NOT:", !aBoolean)            //fals

	//Numerics. Integers
	//int8, int16, int32, int64, int
	//uint8, uint16, uint32, uint64, uint
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	// Value of a: 32 Value of b: 92 Sum of a+b: 124

	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", a) // Type is int

	//Узнаем, сколько байт занимает переемнная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))
	// Type int size of 8 bytes

	//Эксперимент. При использовании короткого объявления - тип для целого числа - int платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))
	// Type int size of 8 bytes

	//Эксперимент 2. Используйте явное приведение типов при необходимости если уверены что не произойдет коллиззии
	var first32 int32 = 12
	var second64 int64 = 13
	//fmt.Println(first32 + second64)
	// будет ошибка

	fmt.Println(int64(first32) + second64) // 25

	//Эксперимент 3. Если проводятся арифметические операции
	// над int и intX , то обязательно нужно использовать механизм приведения. Т.к. int != int64
	var third64 int64 = 16123414
	var fourthInt int = 156234

	//fmt.Println(third64 + fourthInt) // ошибка

	fmt.Println(third64 + int64(fourthInt)) //16279648

	// Аналогичным образом утсроены unit8, uint16, uint32, uint64, uint

	//Numerics. Float
	//float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a %T and type of b %T\n", floatFirst, floatSecond)
	// type of a float64 and type of b float64

	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)           // Sum: 18.21 and Sub: -6.8699999999999
	fmt.Printf("Sum: %.3f and Sub: %.3f", sum, sub) // Sum: 18.21 and Sub: -6.870
	// Важно использовать Printf() при float

	//Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2) // (17 + 39i)
	fmt.Println(c1 * c2) // (-164 + 244i)
	//странное решение по формуле (a + bi)*(c + di)=(ac - bd)+(ad + bc)i
	//(5 + 7)*(12 + 32i)=(5 * 12 - 7 * 32)+(5 * 32 + 7 * 12)i

	//Strings. Строка - это набор БАЙТ
	name := "Федя"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat) // Full name: Федя Pupkin

	fmt.Println("Length of string :", name, len(name)) // Функция len() возвращает количество элементов в наборе
	// Length of string : Федя 8 (считает количество байт а не символов)
	name = "Fedya"
	fmt.Println("Length of string :", name, len(name))
	// Length of string : Fedya 5

	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name)) // Amount of chars: Федя 4

	//Rune - руна. Это один utf-ный символ.
	//Поиск подстроки в строке
	totalString, subString := "ABCDEDFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString)) //true

	subString = "abc"
	fmt.Println(strings.Contains(totalString, subString)) // false

	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' // Для инициализации руны символьным значением - используйте ''
	var thirdRune rune = 234
	fmt.Println(sampleRune) // 0 (это int32 по дефолтному состоянию)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune) // Rune as char Q
	fmt.Printf("Rune as char %c\n", thirdRune)   // Rune as char ё

	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "a")) // -1 if first < second, 0 if first == second, 1 if first > second

	var aByte byte              // alias uint8
	fmt.Println("Byte:", aByte) // Byte: 0

}

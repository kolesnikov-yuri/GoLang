//задача H//////////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		width  int
// 		height int
// 	)

// 	fmt.Scan(&width, &height)

// 	var perimeter = (width + height) * (width + height)

// 	fmt.Print(perimeter)

// }

//задача I //////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		aaa string
// 		bbb string
// 		ccc string
// 	)

// 	fmt.Scan(&aaa, &bbb, &ccc)

// 	var perimeter = (ccc + bbb + aaa)

// 	fmt.Print(perimeter)

// }

//задача J //////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		aaa float32
// 		bbb float32
// 	)

// 	fmt.Scan(&aaa, &bbb)

// 	var perimeter = (bbb + aaa)

// 	if int(perimeter)%2 == 0 {
// 		fmt.Print("ЧЁТНОЕ")
// 	} else {
// 		fmt.Print("НЕЧЁТНОЕ")
// 	}

// }

//задача K //////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	var text string
// 	var context string
// 	context = "чат"

// 	fmt.Scan(&text)
// 	if strings.Contains(text, context) {
// 		fmt.Println("БОТ")
// 	} else {
// 		fmt.Println("НЕ БОТ")
// 	}

// }

//задача L //////////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode/utf8"
// )

// func main() {

// 	var login string
// 	var pochta string

// 	fmt.Scan(&login, &pochta)

// 	if utf8.RuneCountInString(login) < 10 || strings.Contains(login, "@") {
// 		fmt.Println("Некорректный логин")
// 	} else if strings.Contains(pochta, "@") && strings.Contains(pochta, ".") {
// 		fmt.Println("ОК")
// 	} else {
// 		fmt.Println("Некорректная почта")
// 	}

// }

//задача M //////////////////////////////////////

// package main

// import "fmt"

// func main() {

// 	var (
// 		check_1,
// 		check_2,
// 		check_3 string
// 	)
// 	fmt.Scan(&check_1, &check_2, &check_3)

// 	if (check_1 == "раз" || check_1 == "один") && check_2 == "два" && check_3 == "три" {
// 		fmt.Printf("ОК")
// 	} else if check_1 == "1" && check_2 == "2" && check_3 == "3" {
// 		fmt.Printf("ОК")
// 	} else {
// 		fmt.Printf("НЕ ПРАВИЛЬНО")
// 	}
// }

//задача N //////////////////////////////////////

// package main

// import "fmt"

// var (
// 	X_1,
// 	Y_1,
// 	X_2,
// 	Y_2 int
// )

// func main() {

// 	fmt.Scan(&X_1, &Y_1, &X_2, &Y_2)

// 	fmt.Println(X_1,
// 		Y_1,
// 		X_2,
// 		Y_2)
// }
/////////////Задача квадратного уравнения//////////////////////
//  Условие
/*
Задача "написать код на Go, решающий квадратное уравнение ax² + bx + c = 0"
включает в себя ввод коэффициентов a, b, c, вычисление дискриминанта D = b² - 4ac,
 определение количества корней по знаку D и
 расчет самих корней по формулам x1 = (-b + √D) / 2a и x2 = (-b - √D) / 2a для D ≥ 0.
*/

///////////////////////Ответ///////////////////////////////////////

/*

package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	D := b*b - 4*a*c

	if D < 0 {
		fmt.Println("Корней нет")
	}
	if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("Один корень: x = %.2f", x)
	}
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("Два корня: x1 = %.2f, x2 = %.2f", x1, x2)
	}
}
*/

///////////////////////////////////////////////
//  Ответ улучшенный

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("Бесконечно много корней")
			} else {
				fmt.Println("Корней нет")
			}
		} else {
			x := -c / b
			fmt.Printf("Линейное уравнение, корень: x = %.2f\n", x)
		}
		return
	}

	D := b*b - 4*a*c

	if D < 0 {
		fmt.Println("Корней нет")
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("Один корень: x = %.2f\n", x)
	} else {
		sqrtD := math.Sqrt(D)
		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		fmt.Printf("Два корня: x1 = %.2f, x2 = %.2f\n", x1, x2)
	}
}

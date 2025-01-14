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

package main

import "fmt"

func main() {

	var (
		check_1,
		check_2,
		check_3 string
	)
	fmt.Scan(&check_1, &check_2, &check_3)

	if (check_1 == "раз" || check_1 == "один") && check_2 == "два" && check_3 == "три" {
		fmt.Printf("ОК")
	} else if check_1 == "1" && check_2 == "2" && check_3 == "3" {
		fmt.Printf("ОК")
	} else {
		fmt.Printf("НЕ ПРАВИЛЬНО")
	}

}

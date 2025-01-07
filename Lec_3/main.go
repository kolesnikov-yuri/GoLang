// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	var (
// 		age  int
// 		name string
// 	)

// 	// fmt.Scan(&age)
// 	// fmt.Scan(&name)
// 	fmt.Scan(&age, &name)

// 	fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

// 	//Для ручного использования потока ввода
// 	fmt.Fscan(os.Stdin, &age)
// 	fmt.Println("New age:", age)

// }

/////////////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		name   string
// 		famili string
// 		age    int
// 	)

// 	fmt.Scan(&name, &famili, &age)

// 	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", name, famili, age)

// }
////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		drink   string
// 		meal    string
// 		subMeal string
// 		time    string
// 	)

// 	fmt.Scan(&drink, &meal, &subMeal, &time)

// 	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)

// }

//////////////////////////////////////////////////////////////////////////////////////

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var (
// 		world_1 string
// 		world_2 string
// 		world_3 string
// 		world_4 string
// 		world_5 string
// 	)

// 	fmt.Scan(&world_1, &world_2, &world_3, &world_4, &world_5)

// 	fmt.Printf("%s - %s\n", world_4, world_5)
// 	fmt.Printf("%s - %s\n", world_3, world_5)
// 	fmt.Printf("%s - %s\n", world_2, world_5)
// 	fmt.Printf("%s - %s\n", world_1, world_5)
// }

/////////////////////////////////////////////////////

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

// 	var perimeter = (width + height) * 2
// 	var area = width * height

// 	fmt.Printf("Периметр: %d", perimeter)
// 	fmt.Printf("\nПлощадь: %d", area)

// }

////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {

	var (
		width  int
		height int
	)

	fmt.Scan(&width, &height)

	var perimeter = (width + height) * (width + height)

	fmt.Print(perimeter)

}

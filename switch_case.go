package main

import "fmt"

func main() {

	// switch statement (switch case default)
	// switch expression {
	// case value1:
	// 	code to be executed if expression equals value1
	//fallthrough
	// case value2:
	// 	code to be executed if expression equals value2
	// case value3:
	// 	code to be executed if expression equals value3
	// default:
	// 	code to be executed if expression does not match any value
	// }

	// switch statement (switch case break default)
	// switch expression {
	// case value1:
	// 	code to be executed if expression equals value1
	// break
	// case value2:
	// 	code to be executed if expression equals value2
	// break
	// case value3:
	// 	code to be executed if expression equals value3
	//break
	// default:
	// 	code to be executed if expression does not match any value
	// }

	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")
	// case "banana":
	// 	fmt.Println("It's and banana.")
	// default:
	// 	fmt.Println("Unkown fruit!")
	// }

	// //Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}

}

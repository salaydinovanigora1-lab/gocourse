package main

import "go/constant"

type EmployeeGoogle struct {
	FirstName string
	Lastname string
	age int
}

type EmployeeApple struct {
	FirstName string
	Lastname string
	age int
}

func main() {
	//PascalCase
	//Eg. CalculateArea, Userinfo, NewHTTPRequest
	//structs, interfaces, enums

	//snake_case
	//Eg. user_id, first_name, http_request

	//UPPERCASE
	Use case is constants

	//mixedCase
	//Eg.javaScript, htmlDocument, isValid

	const MAXRETRIES =5
	var employeeID =1001
	fmt.Println("Employee ID:", employeeID)
}
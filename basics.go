package main

//3. IMPORT STATEMENT: Bringing in the tool to print text
import "fmt"

// 6. CONSTANT: A value that never changes
const AppName = "MyCoolApp"

func main() {
	//4 & 5. VARIABLES & DATA TYPES:
	//'username' is string variable, 'score' is an int variable
	var username string = "Alice"
	var score int = 100 //Go figures out this is an int automatically

	//Using the Standard Library (fmt) to print it out
	fmt.Println(AppName, "welcomes", username, "with a score of", score)
}

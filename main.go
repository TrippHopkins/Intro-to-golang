package main

import "fmt"

//declares what file this is. GoLang uses different
//code for each package.
//main is a special package which configures the exicutables

// "fmt" is a liberary that includes the "printing to terminal" functions

func main() {

	var Num1 float32 = 20.5
	var Num2 float32 = .5
	var subtractResult float32 = subtract(Num1, Num2)
	fmt.Println(subtractResult)

}
func subtract(Num1 float32, Num2 float32) float32 {

	var subtractionResult float32 = Num1 - Num2

	return subtractionResult
}

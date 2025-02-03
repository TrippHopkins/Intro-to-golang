package main

import (
	"errors"
	"fmt"
)

//declares what file this is. GoLang uses different
//code for each package.
//main is a special package which configures the exicutables

// "fmt" is a liberary that includes the "printing to terminal" functions

func main() {

	var Num1 int = 20
	var Num2 int = 0
	var result, remainder, err = divide(Num1, Num2)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v", result)
	default:
		fmt.Printf("The result is %v and the remainder is %v", result, remainder)

	}
	/*if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result is %v", result)
	} else {
		fmt.Printf("The result is %v and the remainder is %v", result, remainder)

	}*/

}

func divide(Num1 int, Num2 int) (int, int, error) {
	var err error
	if Num2 == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = Num1 / Num2
	var remainder int = Num1 % Num2

	return result, remainder, err
}

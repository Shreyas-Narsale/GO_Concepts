package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func main() {

	/*
		error
			error is a built-in interface .
			type error interface {
		    	Error() string
			}


		errors
			errors is a package

	*/

	//Conversion
	var err error
	err = fmt.Errorf("hello")
	fmt.Println("error ", err.Error())

	var custom_err CustomError
	custom_err.Code = 100
	custom_err.Message = "user not found"

	var err1 error
	err1 = &custom_err
	fmt.Println(err1.Error())

}

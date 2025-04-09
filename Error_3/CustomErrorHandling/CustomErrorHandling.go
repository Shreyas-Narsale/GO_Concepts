package CustomErrorHandling

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("erroCode is: %d error msg is:%s", ce.Code, ce.Message)
}

func CustomErrorHandling() {

	var customerror CustomError
	customerror.Code = 400
	customerror.Message = "Bad Request"

	var err error
	err = &customerror
	fmt.Println(err.Error())
}

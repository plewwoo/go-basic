package helper

import "fmt"

func PanicIfError(err interface{}) {
	if err != nil {
		fmt.Println(err)
	}
}

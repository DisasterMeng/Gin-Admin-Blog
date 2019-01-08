package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		//logging.Info(err.Key, err.Message)
		fmt.Println(err.Key, err.Message)
	}

	return
}

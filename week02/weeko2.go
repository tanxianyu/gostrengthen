package main

import (
	"fmt"
	"gostrengthen/week02/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {

	err, dataList := service.Select()
	if err != nil {
		fmt.Printf("original error:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
	fmt.Println(dataList)
}

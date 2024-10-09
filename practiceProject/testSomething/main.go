package main

import (
	"errors"
	"fmt"
)

func someFunction() error {
	// 返回一个错误
	return fmt.Errorf("this is an error")
}

func main() {
	//err := gorm.ErrRecordNotFound
	//err := errors.New("this is an error")
	//err := someFunction()
	//if err != nil {
	//	t := reflect.TypeOf(err).String()
	//	fmt.Println("Error type is", t)
	//	fmt.Print(err)
	//}
	err1 := errors.New("this is an error")
	err2 := errors.New("this is an error")
	if errors.Is(err1, err2) {
		fmt.Println("err1 == err2")
	}

	if err1.Error() == err2.Error() {
		fmt.Println("yes")
	}

}

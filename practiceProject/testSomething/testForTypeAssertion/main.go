package main

import "fmt"

// MySQLError is an error type which represents a single MySQL error
type MySQLError struct {
	Number   uint16
	SQLState [5]byte
	Message  string
}

func (me *MySQLError) Error() string {
	if me.SQLState != [5]byte{} {
		return fmt.Sprintf("Error %d (%s): %s", me.Number, me.SQLState, me.Message)
	}

	return fmt.Sprintf("Error %d: %s", me.Number, me.Message)
}

func (me *MySQLError) Is(err error) bool {
	if merr, ok := err.(*MySQLError); ok {
		return merr.Number == me.Number
	}
	return false
}

func main() {
	var err error
	mysqlerr := &MySQLError{1, [5]byte{'d', 'a', 'a', 'a', 'a'}, "abcde"}
	err = mysqlerr

	if errN, ok := err.(*MySQLError); ok {
		fmt.Println(errN)
	}
	if number, ok := err.(uint16); ok {
		fmt.Println(number)
	}
}

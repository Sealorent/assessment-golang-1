package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/tes1/v2/student"
)

var students []student.Student

func init() {
	students = student.SliceOfStudent()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one argument.")
		return
	}

	argument := os.Args[1]

	intValue, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if intValue < 1 || intValue > len(students) {
		fmt.Println("Error: ", "ID not found")
		return
	}

	for _, v := range students {
		if v.Id == intValue {
			v.InvokeGreetings()
			return
		}
	}
	fmt.Println("Error: ", "ID not found")
}

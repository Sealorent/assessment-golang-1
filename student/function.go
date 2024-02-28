package student

import "fmt"

func (s Student) GetId() {
	fmt.Println("ID :", s.Id)
}

func (s Student) GetName() {
	fmt.Println("Name :", s.Name)
}

func (s Student) GetAdress() {
	fmt.Println("Adress :", s.Adress)
}

func (s Student) GetJob() {
	fmt.Println("Job :", s.Job)
}

func (s Student) GetReason() {
	fmt.Println("Reason :", s.Reason)
}

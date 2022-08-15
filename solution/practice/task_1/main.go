package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
}

func (h *Human) Init(name, surname string) {
	h.name = name
	h.surname = surname
	h.age = 0
}

func (h *Human) ChangeName(newName string) {
	h.name = newName
}

func (h *Human) ChangeSurname(newSurname string) {
	h.surname = newSurname
}

func (h *Human) CelebrateBirthday() {
	h.age++
}

type Action struct {
	Human
}

func (act *Action) Presentate() string {
	return fmt.Sprintf("Name: %s\nSurname: %s\nAge: %d", act.name, act.surname, act.age)
}

func main() {
	humanAction := Action{}
	humanAction.Init("Ivan", "Ivanov")
	fmt.Println(humanAction.Presentate())
}

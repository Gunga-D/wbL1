package main

import "fmt"

type USB interface {
	ConnectByUSB()
}

type TypeC struct {
}

func (con *TypeC) ConnectByUSB() {
	fmt.Println("Connected by type c")
}

type Lightning struct {
}

func (con *Lightning) Connect() {
	fmt.Println("Connected by lightning")
}

type AdapterFromTypeCToLightning struct {
	lightning *Lightning
}

func (adapt *AdapterFromTypeCToLightning) ConnectByUSB() {
	adapt.lightning.Connect()
}

type Phone struct {
}

func (p *Phone) Charge(conn USB) {
	conn.ConnectByUSB()
}

func main() {
	typec := &TypeC{}
	lightning := &Lightning{}
	adapter := &AdapterFromTypeCToLightning{lightning: lightning}

	phone := Phone{}
	phone.Charge(typec)
	fmt.Println("OR")
	phone.Charge(adapter)
}

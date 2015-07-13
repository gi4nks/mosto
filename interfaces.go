package main


type Animal interface {
	Greet() string	
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

type Bird struct {
	Animal
}

func (c Cat) Greet() string {
	return "MiaoMiao"
}

func (d Dog) Greet() string {
	return "BauBau"
}

func (b Bird) Greet() string {
	return "PioPio"
}
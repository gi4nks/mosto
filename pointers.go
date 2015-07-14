package main

type Sample struct {
	privateField string
	PublicField string
}


func (s Sample) String() string {
	return "{ privateField: " + s.privateField + ", publicField: " + s.PublicField + "}"
}

func (s Sample) Stateless() {
	s.privateField = "Stateless - changed"
	s.PublicField = "Stateless - changed"	
	
	tracer.Warning("Inside Stateless() function --> " + s.String())
}

func (s *Sample) Stateful() {
	s.privateField = "Stateful - changed"
	s.PublicField = "Stateful - changed"	
	
	tracer.Warning("Inside Stateful() function --> " + s.String())
}



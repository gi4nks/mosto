package main

import ( 
	"errors"
	"fmt"
	
)

//type write_1 func(string, ...interface {})
type write_1 func(string, ...interface {})
type write_3 func(string, ...interface {}) (int, error)
type write_2 func()

func trace_3(w write_3, s string) error {
	if (s == "") {
		return  errors.New("s is emtpy")
	}
	w(s)
	return nil
}

func trace_2(w write_2) error {
	w()
	return nil
}

/*
func plusTwo() (func(v int) (int)) {
    return func(v int) (int) {
        return v+2
    }
}

func plusX(x int) (func(v int) (int)) {
    return func(v int) (int) {
        return v+x
    }
}
*/
func Functions() {
  	
	trace_3(fmt.Printf, "trace_3") 
	
	trace_2(func() { fmt.Printf("trace_2") }) 
	
	/*
	p := plusTwo()
    fmt.Printf("3+2: %d\n", p(3))

    px := plusX(3)
    fmt.Printf("3+3: %d\n", px(3))
	*/
}
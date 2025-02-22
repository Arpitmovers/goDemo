package main

import (
	"fmt"
	"log"
)

type customErr struct {
	lat  string
	long string
	err  error
}

// implement Error interface
func (se customErr) Error() string {
	return fmt.Sprintf("math error : %v %v  %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-11)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {

	if f < 0 {

		// way 1: using errors.New
		//	a := customErr{lat: "50.", long: "long", err: errors.New("more coffee needed")}
		//return 0, a

		// way 2 : using fmt.Errorf
		e := fmt.Errorf("more coffee needed -value was %v", f)
		return 0, customErr{"lat 41.1", "long 55.", e}

	}
	return f, nil
}

package main

import (
	"fmt"
)

type Employee interface {
	salary() int
}
type FullTime struct {
	days int
}
type Contractor struct {
	days int
}
type Freelancer struct {
	hours int
}

func (x FullTime) salary() int {
	return 500*x.days + 1000
}

func (x Contractor) salary() int {
	return 100*x.days + 200
}

func (x Freelancer) salary() int {
	return 10*x.hours + 1800
}

func main() {
	x := FullTime{days: 20}
	y := Contractor{days: 20}
	z := Freelancer{hours: 150}
	fmt.Println(x.salary())
	fmt.Println(y.salary())
	fmt.Println(z.salary())
}

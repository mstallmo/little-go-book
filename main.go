package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func main() {
	goku := NewSaiyan("goku", 9000)
	goku.Super()
	fmt.Println(goku.Power)
}

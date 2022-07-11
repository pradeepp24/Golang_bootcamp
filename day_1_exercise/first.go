package main

import (
	"fmt"
	"strconv"
)

type Matrix struct {
	rows    int `json:"rows"`
	columns int
	//elements := make([][]int, rows)
	elements [][]int
}

func (x *Matrix) intialize(rows, columns int) {
	x.rows = rows
	x.columns = columns
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, columns)
	}
	x.elements = elements
}
func (x *Matrix) getRows() int {
	return x.rows
}
func (x *Matrix) getColumns() int {
	return x.columns
}
func (x *Matrix) setValues(i, j, value int) {
	x.elements[i][j] = value
}
func (x *Matrix) addTwoMatrix(y Matrix) {
	if (x.rows == y.rows) && (x.columns == y.columns) {
		for i := 0; i < x.rows; i++ {
			for j := 0; j < x.columns; j++ {
				x.setValues(i, j, x.elements[i][j]+y.elements[i][j])
			}
		}
	}
}
func (x *Matrix) printJson() string {
	s := "["
	for i := 0; i < x.rows; i++ {
		s += "["
		for j := 0; j < x.columns; j++ {
			temp := strconv.Itoa(x.elements[i][j])
			s += temp
			if j != x.columns-1 {
				s += ", "
			}
		}
		s += "]"
	}
	s += "]"

	return fmt.Sprintf("{Rows:%d, Columns:%d ,elements:%s }", x.rows, x.columns, s)
}

func (x *Matrix) printValues() {
	for i := 0; i < x.rows; i++ {
		for j := 0; j < x.columns; j++ {
			fmt.Print(x.elements[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	m := Matrix{}
	m.intialize(5, 7)
	ROW := 5
	COLUMN := 7
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			m.setValues(i, j, i+26*j)
		}
	}
	l := Matrix{}
	l.intialize(ROW, COLUMN)
	for i := 0; i < ROW; i++ {
		for j := 0; j < COLUMN; j++ {
			l.setValues(i, j, i+j)
		}
	}
	fmt.Println(m.getRows())
	fmt.Println(m.getColumns())

	m.printValues()
	l.printValues()

	m.addTwoMatrix(l)
	
	fmt.Println(m.printJson())

}

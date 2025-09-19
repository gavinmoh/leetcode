package main

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	Cells [][]int
}

func Constructor(rows int) Spreadsheet {
	cells := make([][]int, rows)
	for row := 0; row < rows; row++ {
		cells[row] = make([]int, 26)
	}

	return Spreadsheet{Cells: cells}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	row, col := GetRowColFromCellStr(cell)
	this.Cells[row][col] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	row, col := GetRowColFromCellStr(cell)
	this.Cells[row][col] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	operands := strings.Split(formula[1:], "+")

	return this.GetOperandValue(operands[0]) + this.GetOperandValue(operands[1])
}

func GetRowColFromCellStr(cellStr string) (int, int) {
	col := int(rune(cellStr[0]) - 'A')
	row, _ := strconv.Atoi(cellStr[1:])

	// row is given as 1-indexed
	return row - 1, col
}

func (this *Spreadsheet) GetOperandValue(str string) int {
	if rune(str[0]) >= 'A' && rune(str[1]) <= 'Z' {
		row, col := GetRowColFromCellStr(str)
		return this.Cells[row][col]
	}

	val, _ := strconv.Atoi(str)
	return val
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */

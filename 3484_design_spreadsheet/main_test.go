package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpreadsheet(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rows := 3

		spreadsheet := Constructor(rows)
		assert.Equal(t, 12, spreadsheet.GetValue("=5+7"))

		spreadsheet.SetCell("A1", 10)
		assert.Equal(t, 16, spreadsheet.GetValue("=A1+6"))

		spreadsheet.SetCell("B2", 15)
		assert.Equal(t, 25, spreadsheet.GetValue("=A1+B2"))

		spreadsheet.ResetCell("A1")
		assert.Equal(t, 15, spreadsheet.GetValue("=A1+B2"))
	})

	t.Run("test case 2", func(t *testing.T) {
		rows := 24

		spreadsheet := Constructor(rows)
		spreadsheet.SetCell("B24", 66688)
		spreadsheet.ResetCell("O15")
	})
}

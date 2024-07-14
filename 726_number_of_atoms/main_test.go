package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOfAtoms(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		formula := "H2O"
		expected := "H2O"

		assert.Equal(t, expected, countOfAtoms(formula))
	})

	t.Run("test case 2", func(t *testing.T) {
		formula := "Mg(OH)2"
		expected := "H2MgO2"

		assert.Equal(t, expected, countOfAtoms(formula))
	})

	t.Run("test case 3", func(t *testing.T) {
		formula := "K4(ON(SO3)2)2"
		expected := "K4N2O14S4"

		assert.Equal(t, expected, countOfAtoms(formula))
	})

	t.Run("test case 4", func(t *testing.T) {
		formula := "((N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"
		expected := "B18900Be18984C4200H5446He1386Li33894N50106O22638"

		assert.Equal(t, expected, countOfAtoms(formula))
	})
}

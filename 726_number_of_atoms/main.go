package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Atom struct {
	Symbol        string
	Count         int
	SymbolLetters []byte
	CountLetters  []byte
}

func countOfAtoms(formula string) string {
	i := 0
	atomMap := make(map[string]int)

	for i < len(formula) {
		atoms, next := parse(formula[i:])
		for _, atom := range atoms {
			atomMap[atom.Symbol] += atom.Count
		}
		i = next
	}

	symbols := make([]string, 0, len(atomMap))
	for symbol := range atomMap {
		symbols = append(symbols, symbol)
	}
	sort.Strings(sort.StringSlice(symbols))

	result := ""
	for _, symbol := range symbols {
		result += symbol
		if atomMap[symbol] > 1 {
			result += fmt.Sprintf("%d", atomMap[symbol])
		}
	}

	return result
}

func stringToAtoms(formula []byte) []*Atom {
	atoms := []*Atom{}
	for _, letter := range formula {
		if letter >= 'A' && letter <= 'Z' {
			atom := &Atom{
				SymbolLetters: []byte{letter},
				CountLetters:  []byte{},
				Count:         1,
			}
			atoms = append(atoms, atom)
			continue
		}

		last := atoms[len(atoms)-1]
		if letter >= 'a' && letter <= 'z' {
			last.SymbolLetters = append(last.SymbolLetters, letter)
		} else {
			last.CountLetters = append(last.CountLetters, letter)
		}
	}
	for _, atom := range atoms {
		// convert to symbol & count
		atom.Symbol = string(atom.SymbolLetters)
		if len(atom.CountLetters) > 0 {
			atom.Count, _ = strconv.Atoi(string(atom.CountLetters))
		}
	}

	return atoms
}

// taking a string and extract the formula within the next ')'
// return a slice of atoms and the index after the ')' or the multiplier
func parse(str string) ([]*Atom, int) {
	multiplier := []byte{}
	formula := []byte{}
	close := -1
	i := 0
	atoms := []*Atom{}
	for i < len(str) {
		letter := str[i]

		if close != -1 && !(letter >= '0' && letter <= '9') {
			break
		}

		if letter == '(' {
			nestedAtoms, next := parse(str[i+1:])
			atoms = append(atoms, nestedAtoms...)
			i += next
			continue
		}

		if letter == ')' {
			close = i
		} else if close != -1 && letter >= '0' && letter <= '9' {
			multiplier = append(multiplier, letter)
		} else {
			formula = append(formula, letter)
		}
		i++
	}

	// take formula and convert it to slice of atoms
	atoms = append(atoms, stringToAtoms(formula)...)

	if len(multiplier) > 0 {
		intMultiplier, _ := strconv.Atoi(string(multiplier))
		for _, atom := range atoms {
			atom.Count *= intMultiplier
		}
	}

	return atoms, i + 1
}

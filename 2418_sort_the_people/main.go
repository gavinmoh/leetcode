package main

import "sort"

type Person struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	people := []*Person{}
	for i := 0; i < len(names); i++ {
		people = append(people, &Person{Name: names[i], Height: heights[i]})
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})

	for i, person := range people {
		names[i] = person.Name
	}

	return names
}

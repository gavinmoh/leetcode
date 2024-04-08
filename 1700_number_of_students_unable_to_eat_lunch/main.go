package main

func countStudents(students []int, sandwiches []int) int {
	count := 0
	for count < len(students) {
		if students[0] == sandwiches[0] {
			// remove both student and sandwitch
			students = students[1:]
			sandwiches = sandwiches[1:]
			// reset loop counter
			count = 0
		} else {
			// student moved to the end
			students = append(students[1:], students[0])
			count++
		}
	}

	return len(students)
}

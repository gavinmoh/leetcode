package main

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Top() T {
	return s.data[s.Size()-1]
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() T {
	top := s.Top()
	s.data = s.data[:s.Size()-1]
	return top
}

func (s *Stack[T]) NotEmpty() bool {
	return s.Size() > 0
}

type Pair struct {
	Start  int
	Height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := Stack[Pair]{}
	maxArea := 0

	for i, height := range heights {
		start := i
		for stack.NotEmpty() && stack.Top().Height > height {
			pair := stack.Pop()
			index, height := pair.Start, pair.Height
			maxArea = max(maxArea, height*(i-index))
			start = index
		}
		stack.Push(Pair{Start: start, Height: height})
	}

	for stack.NotEmpty() {
		pair := stack.Pop()
		index, height := pair.Start, pair.Height
		maxArea = max(maxArea, height*(n-index))
	}

	return maxArea
}

func maximalRectangle(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0
	var previousHeights []int

	for row := 0; row < rows; row++ {
		heights := make([]int, cols)
		for i, char := range matrix[row] {
			height := 0
			if char == '1' {
				height += 1
				if row > 0 {
					height += previousHeights[i]
				}
			}
			heights[i] = height
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
		previousHeights = heights
	}

	return maxArea
}

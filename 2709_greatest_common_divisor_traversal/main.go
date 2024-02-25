package main

type UnionFind struct {
	parents []int
	sizes   []int
	count   int
}

func Constructor(n int) *UnionFind {
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		sizes[i] = 1
	}
	return &UnionFind{
		parents: parents,
		sizes:   sizes,
		count:   n,
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.find(uf.parents[x])
	}
	return uf.parents[x]
}

func (uf *UnionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return
	}
	if uf.sizes[px] < uf.sizes[py] {
		uf.parents[px] = py
		uf.sizes[py] += uf.sizes[px]
	} else {
		uf.parents[py] = px
		uf.sizes[px] += uf.sizes[py]
	}
	uf.count -= 1
}

func canTraverseAllPairs(nums []int) bool {
	uf := Constructor(len(nums))
	factorIndex := make(map[int]int)

	for i, num := range nums {
		f := 2
		for f*f <= num {
			if num%f == 0 {
				if _, ok := factorIndex[f]; ok {
					uf.union(i, factorIndex[f])
				} else {
					factorIndex[f] = i
				}
				for num%f == 0 {
					num = num / f
				}
			}
			f += 1
		}
		if num > 1 {
			if _, ok := factorIndex[num]; ok {
				uf.union(i, factorIndex[num])
			} else {
				factorIndex[num] = i
			}
		}
	}

	return uf.count == 1
}

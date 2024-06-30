package main

type UF struct {
	N      int
	Parent []int
	Rank   []int
}

func (uf *UF) Find(i int) int {
	if uf.Parent[i] != i {
		uf.Parent[i] = uf.Find(uf.Parent[i])
	}
	return uf.Parent[i]
}

func (uf *UF) Union(i, j int) int {
	p1, p2 := uf.Find(i), uf.Find(j)

	if p1 == p2 {
		return 0
	}

	if uf.Rank[p1] > uf.Rank[p2] {
		uf.Parent[p2] = p1
		uf.Rank[p1] += uf.Rank[p2]
	} else {
		uf.Parent[p1] = p2
		uf.Rank[p2] += uf.Rank[p1]
	}
	uf.N--
	return 1
}

func (uf *UF) IsConnected() bool { return uf.N == 1 }

func Constructor(n int) *UF {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 0; i <= n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	uf := &UF{
		N:      n,
		Parent: parent,
		Rank:   rank,
	}
	return uf
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice, bob := Constructor(n), Constructor(n)
	count := 0 // number of edges we have to keep

	for _, edge := range edges {
		t, src, dst := edge[0], edge[1], edge[2]

		if t == 3 {
			count += (alice.Union(src, dst) | bob.Union(src, dst))
		}
	}

	for _, edge := range edges {
		t, src, dst := edge[0], edge[1], edge[2]

		if t == 1 {
			count += alice.Union(src, dst)
		} else if t == 2 {
			count += bob.Union(src, dst)
		}
	}

	if bob.IsConnected() && alice.IsConnected() {
		return len(edges) - count
	}

	return -1
}

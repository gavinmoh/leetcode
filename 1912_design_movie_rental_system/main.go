package main

import "container/heap"

type Entry struct {
	Shop     int
	Movie    int
	Price    int
	Index    int
	IsRented bool
}

type EntryIndex struct {
	Shop  int
	Movie int
}

type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].IsRented {
		if pq[i].Price == pq[j].Price {
			if pq[i].Shop == pq[j].Shop {
				return pq[i].Movie < pq[j].Movie
			}
			return pq[i].Shop < pq[j].Shop
		}
	} else {
		if pq[i].Price == pq[j].Price {
			return pq[i].Shop < pq[j].Shop
		}
	}

	return pq[i].Price < pq[j].Price
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	entry := x.(*Entry)
	entry.Index = n
	*pq = append(*pq, entry)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	entry := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	entry.Index = -1
	*pq = old[0 : n-1]
	return entry
}

type MovieRentingSystem struct {
	RentedMovies    *PriorityQueue
	RentedMoviesMap map[EntryIndex]*Entry
	Shops           []map[int]*Entry
	Movies          map[int]*PriorityQueue
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	rentedMovies := PriorityQueue{}
	heap.Init(&rentedMovies)
	rentedMoviesMap := make(map[EntryIndex]*Entry)

	shops := make([]map[int]*Entry, n)
	for i := 0; i < n; i++ {
		shops[i] = map[int]*Entry{}
	}
	movies := make(map[int]*PriorityQueue)

	for _, e := range entries {
		entry := Entry{Shop: e[0], Movie: e[1], Price: e[2]}
		shops[entry.Shop][entry.Movie] = &entry

		if _, ok := movies[entry.Movie]; !ok {
			pq := PriorityQueue{}
			heap.Init(&pq)
			movies[entry.Movie] = &pq
		}

		heap.Push(movies[entry.Movie], &entry)
	}

	return MovieRentingSystem{
		RentedMovies:    &rentedMovies,
		RentedMoviesMap: rentedMoviesMap,
		Shops:           shops,
		Movies:          movies,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	if _, ok := this.Movies[movie]; !ok {
		return []int{}
	}

	result := []int{}
	popped := []*Entry{}
	for i := 0; i < 5 && this.Movies[movie].Len() > 0; i++ {
		entry := heap.Pop(this.Movies[movie]).(*Entry)
		popped = append(popped, entry)
	}

	for _, entry := range popped {
		result = append(result, entry.Shop)
		heap.Push(this.Movies[movie], entry)
	}

	return result
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	entry := this.Shops[shop][movie]

	// remove from shop and movies list
	delete(this.Shops[shop], movie)
	heap.Remove(this.Movies[movie], entry.Index)

	// mark it as rented
	entry.IsRented = true

	// place it in rented movie
	heap.Push(this.RentedMovies, entry)
	entryIndex := EntryIndex{Shop: entry.Shop, Movie: entry.Movie}
	this.RentedMoviesMap[entryIndex] = entry
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	entryIndex := EntryIndex{Shop: shop, Movie: movie}
	entry := this.RentedMoviesMap[entryIndex]

	// remove from rented list
	delete(this.RentedMoviesMap, entryIndex)
	heap.Remove(this.RentedMovies, entry.Index)

	// mark it as not rented
	entry.IsRented = false

	// add back to shop and movies list
	this.Shops[shop][movie] = entry
	heap.Push(this.Movies[movie], entry)
}

func (this *MovieRentingSystem) Report() [][]int {
	result := [][]int{}
	popped := []*Entry{}

	for i := 0; i < 5 && this.RentedMovies.Len() > 0; i++ {
		entry := heap.Pop(this.RentedMovies).(*Entry)
		popped = append(popped, entry)
	}

	for _, entry := range popped {
		result = append(result, []int{entry.Shop, entry.Movie})
		heap.Push(this.RentedMovies, entry)
	}

	return result
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */

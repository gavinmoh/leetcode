package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieRentingSystem(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		entries := [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}}
		sys := Constructor(n, entries)

		assert.Equal(t, []int{1, 0, 2}, sys.Search(1))

		sys.Rent(0, 1)
		sys.Rent(1, 2)
		assert.Equal(t, [][]int{{0, 1}, {1, 2}}, sys.Report())

		sys.Drop(1, 2)
		assert.Equal(t, []int{0, 1}, sys.Search(2))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 10
		entries := [][]int{
			{0, 418, 3}, {9, 5144, 46}, {2, 8986, 29}, {6, 1446, 28}, {3, 8215, 97},
			{7, 9105, 34}, {6, 9105, 30}, {5, 1722, 94}, {9, 528, 40}, {3, 850, 77},
			{3, 7069, 40}, {8, 1997, 42}, {0, 8215, 28}, {7, 4050, 80}, {4, 7100, 97},
			{4, 9686, 32}, {4, 2566, 93}, {2, 8320, 12}, {2, 5495, 56},
		}
		sys := Constructor(n, entries)

		assert.Equal(t, []int{}, sys.Search(7837))
		assert.Equal(t, []int{2}, sys.Search(5495))

		sys.Rent(4, 7100)
		assert.Equal(t, []int{6, 7}, sys.Search(9105))
		assert.Equal(t, []int{6}, sys.Search(1446))
		assert.Equal(t, [][]int{{4, 7100}}, sys.Report())
		assert.Equal(t, []int{}, sys.Search(9869))

		sys.Drop(4, 7100)
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 69
		entries := [][]int{
			{16, 4156, 1511}, {20, 8501, 8417}, {34, 7901, 7776}, {54, 6691, 9511},
			{44, 8931, 8434}, {42, 9640, 5251}, {22, 4534, 9161}, {32, 6506, 6831},
			{13, 8501, 731}, {4, 7610, 8474}, {33, 820, 2341}, {17, 6490, 1161},
			{29, 7120, 2703}, {8, 8723, 7613}, {38, 9544, 1804}, {30, 8723, 1047},
			{1, 5015, 7763}, {60, 1625, 2383}, {29, 3336, 3542}, {39, 7535, 6066},
			{1, 9074, 9400}, {39, 1625, 7944}, {26, 9160, 6874}, {55, 2465, 888},
			{35, 8530, 6025},
		}
		sys := Constructor(n, entries)

		// Operations and assertions
		sys.Rent(32, 6506)
		assert.Equal(t, []int{13, 20}, sys.Search(8501))
		assert.Equal(t, []int{}, sys.Search(6275))
		assert.Equal(t, [][]int{{32, 6506}}, sys.Report())
		sys.Rent(30, 8723)
		sys.Rent(8, 8723)
		assert.Equal(t, [][]int{{30, 8723}, {32, 6506}, {8, 8723}}, sys.Report())
		assert.Equal(t, [][]int{{30, 8723}, {32, 6506}, {8, 8723}}, sys.Report())
		assert.Equal(t, []int{}, sys.Search(6699))
		assert.Equal(t, []int{}, sys.Search(115))
		sys.Rent(20, 8501)
		sys.Rent(16, 4156)
		assert.Equal(t, []int{}, sys.Search(9447))
		sys.Drop(30, 8723)
		sys.Drop(8, 8723)
		sys.Drop(32, 6506)
		sys.Drop(16, 4156)
		sys.Rent(42, 9640)
		assert.Equal(t, [][]int{{42, 9640}, {20, 8501}}, sys.Report())
		assert.Equal(t, [][]int{{42, 9640}, {20, 8501}}, sys.Report())
		sys.Rent(17, 6490)
		sys.Drop(20, 8501)
		assert.Equal(t, []int{}, sys.Search(8175))
		assert.Equal(t, [][]int{{17, 6490}, {42, 9640}}, sys.Report())
		sys.Drop(17, 6490)
		assert.Equal(t, [][]int{{42, 9640}}, sys.Report())
		sys.Drop(42, 9640)
		sys.Rent(54, 6691)
		assert.Equal(t, [][]int{{54, 6691}}, sys.Report())
		assert.Equal(t, []int{60, 39}, sys.Search(1625))
		assert.Equal(t, []int{}, sys.Search(3291))
		sys.Rent(60, 1625)
		sys.Rent(39, 1625)
		assert.Equal(t, [][]int{{60, 1625}, {39, 1625}, {54, 6691}}, sys.Report())
		assert.Equal(t, [][]int{{60, 1625}, {39, 1625}, {54, 6691}}, sys.Report())
		sys.Drop(60, 1625)
		assert.Equal(t, [][]int{{39, 1625}, {54, 6691}}, sys.Report())
		assert.Equal(t, [][]int{{39, 1625}, {54, 6691}}, sys.Report())
		sys.Drop(39, 1625)
		assert.Equal(t, [][]int{{54, 6691}}, sys.Report())
		sys.Drop(54, 6691)
		sys.Rent(8, 8723)
		sys.Drop(8, 8723)
		assert.Equal(t, []int{}, sys.Search(2260))
		sys.Rent(29, 7120)
		assert.Equal(t, []int{}, sys.Search(746))
		sys.Drop(29, 7120)
		sys.Rent(38, 9544)
		sys.Drop(38, 9544)
		assert.Equal(t, [][]int{}, sys.Report())
		sys.Rent(1, 9074)
		sys.Drop(1, 9074)
		sys.Rent(54, 6691)
		sys.Rent(39, 1625)
		sys.Drop(54, 6691)
		assert.Equal(t, [][]int{{39, 1625}}, sys.Report())
		assert.Equal(t, [][]int{{39, 1625}}, sys.Report())
		assert.Equal(t, [][]int{{39, 1625}}, sys.Report())
		assert.Equal(t, [][]int{{39, 1625}}, sys.Report())
		sys.Rent(26, 9160)
		sys.Drop(26, 9160)
		assert.Equal(t, [][]int{{39, 1625}}, sys.Report())
		sys.Drop(39, 1625)
		sys.Rent(42, 9640)
		assert.Equal(t, []int{}, sys.Search(9640))
		sys.Drop(42, 9640)
		assert.Equal(t, [][]int{}, sys.Report())
		sys.Rent(29, 7120)
		assert.Equal(t, []int{}, sys.Search(5630))
		assert.Equal(t, []int{}, sys.Search(1842))
		assert.Equal(t, [][]int{{29, 7120}}, sys.Report())
		sys.Rent(16, 4156)
		assert.Equal(t, [][]int{{16, 4156}, {29, 7120}}, sys.Report())
		assert.Equal(t, [][]int{{16, 4156}, {29, 7120}}, sys.Report())
		sys.Rent(1, 9074)
		assert.Equal(t, [][]int{{16, 4156}, {29, 7120}, {1, 9074}}, sys.Report())
		assert.Equal(t, [][]int{{16, 4156}, {29, 7120}, {1, 9074}}, sys.Report())
		assert.Equal(t, []int{}, sys.Search(7992))
		sys.Rent(4, 7610)
		sys.Rent(29, 3336)
		assert.Equal(t, []int{}, sys.Search(1333))
	})
}

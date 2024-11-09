package main

import "sort"

type Item struct {
	Price  int
	Beauty int
}

type Query struct {
	Index int
	Price int
}

func maximumBeauty(items [][]int, queries []int) []int {
	itemList := make([]*Item, len(items))
	for i, item := range items {
		itemList[i] = &Item{
			Price:  item[0],
			Beauty: item[1],
		}
	}
	sort.SliceStable(itemList, func(i, j int) bool {
		if itemList[i].Price == itemList[j].Price {
			return itemList[i].Beauty < itemList[j].Beauty
		}
		return itemList[i].Price < itemList[j].Price
	})

	queryList := make([]*Query, len(queries))
	for i, price := range queries {
		queryList[i] = &Query{
			Index: i,
			Price: price,
		}
	}
	sort.SliceStable(queryList, func(i, j int) bool {
		return queryList[i].Price < queryList[j].Price
	})

	result := make([]int, len(queries))
	itemIdx := 0
	maxBeauty := 0
	for _, query := range queryList {
		for itemIdx < len(items) && itemList[itemIdx].Price <= query.Price {
			maxBeauty = max(maxBeauty, itemList[itemIdx].Beauty)
			itemIdx++
		}
		result[query.Index] = maxBeauty
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

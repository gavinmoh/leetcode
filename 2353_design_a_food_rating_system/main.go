package main

import (
	"container/heap"
	"strings"
)

type Food struct {
	Name    string
	Rating  int
	Cuisine string
}

type RatingQueue []*Food

func (rq RatingQueue) Len() int {
	return len(rq)
}

func (rq RatingQueue) Less(i, j int) bool {
	if rq[i].Rating == rq[j].Rating {
		return strings.Compare(rq[i].Name, rq[j].Name) < 0
	}
	return rq[i].Rating > rq[j].Rating
}

func (rq RatingQueue) Swap(i, j int) {
	rq[i], rq[j] = rq[j], rq[i]
}

func (rq *RatingQueue) Push(x any) {
	food := x.(*Food)
	*rq = append(*rq, food)
}

func (rq *RatingQueue) Pop() any {
	old := *rq
	n := len(old)
	food := old[n-1]
	old[n-1] = nil
	*rq = old[0 : n-1]
	return food
}

func (rq *RatingQueue) Top() *Food {
	return (*rq)[0]
}

func (rq *RatingQueue) Update(food *Food, rating int) {
	food.Rating = rating
	for i, f := range *rq {
		if f == food {
			heap.Fix(rq, i)
			return
		}
	}
}

type FoodRatings struct {
	queues map[string]*RatingQueue
	foods  map[string]*Food
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	queues := make(map[string]*RatingQueue)
	foodsMap := make(map[string]*Food, len(foods))
	for i, name := range foods {
		cuisine := cuisines[i]
		rating := ratings[i]
		food := &Food{Name: name, Cuisine: cuisine, Rating: rating}
		foodsMap[name] = food
		if _, ok := queues[cuisine]; !ok {
			rq := RatingQueue{}
			heap.Init(&rq)
			queues[cuisine] = &rq
		}
		rq := queues[cuisine]
		heap.Push(rq, food)
	}
	return FoodRatings{queues: queues, foods: foodsMap}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	foodObj := this.foods[food]
	queues := this.queues[foodObj.Cuisine]
	queues.Update(foodObj, newRating)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	rq := this.queues[cuisine]
	return rq.Top().Name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodRatings(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
		cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
		ratings := []int{9, 12, 8, 15, 14, 7}

		foodRatings := Constructor(foods, cuisines, ratings)
		assert.Equal(t, "kimchi", foodRatings.HighestRated("korean"))
		assert.Equal(t, "ramen", foodRatings.HighestRated("japanese"))

		foodRatings.ChangeRating("sushi", 16)
		assert.Equal(t, "sushi", foodRatings.HighestRated("japanese"))

		foodRatings.ChangeRating("ramen", 16)
		assert.Equal(t, "ramen", foodRatings.HighestRated("japanese"))
	})

	t.Run("example 2", func(t *testing.T) {
		foods := []string{"xxdcg", "wfqdeytt", "jqmfm", "ukqbjikyx", "aymciznrnw", "qhjjrvr", "wzcinxg", "ikxj"}
		cuisines := []string{"lruhtqy", "lruhtqy", "lruhtqy", "lruhtqy", "lruhtqy", "lruhtqy", "lruhtqy", "lruhtqy"}
		ratings := []int{8, 6, 1, 17, 20, 2, 17, 14}

		foodRatings := Constructor(foods, cuisines, ratings)
		assert.Equal(t, "aymciznrnw", foodRatings.HighestRated("lruhtqy"))

		foodRatings.ChangeRating("wfqdeytt", 17)
		foodRatings.ChangeRating("aymciznrnw", 9)
		assert.Equal(t, "ukqbjikyx", foodRatings.HighestRated("lruhtqy"))

		foodRatings.ChangeRating("ukqbjikyx", 10)
		assert.Equal(t, "wfqdeytt", foodRatings.HighestRated("lruhtqy"))

		foodRatings.ChangeRating("xxdcg", 11)
		foodRatings.ChangeRating("ikxj", 15)
		foodRatings.ChangeRating("aymciznrnw", 10)
		foodRatings.ChangeRating("wfqdeytt", 10)
		foodRatings.ChangeRating("xxdcg", 16)
		foodRatings.ChangeRating("ikxj", 2)
		foodRatings.ChangeRating("aymciznrnw", 16)
		assert.Equal(t, "wzcinxg", foodRatings.HighestRated("lruhtqy"))

		foodRatings.ChangeRating("wzcinxg", 12)
		assert.Equal(t, "aymciznrnw", foodRatings.HighestRated("lruhtqy"))
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItinerary(t *testing.T) {
	case1 := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	case1Expected := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	assert.Equal(t, case1Expected, findItinerary(case1))

	case2 := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	case2Expected := []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	assert.Equal(t, case2Expected, findItinerary(case2))
}

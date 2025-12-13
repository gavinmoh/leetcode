package main

import "sort"

type Coupon struct {
	Code     string
	Priority int
}

var priorities = map[string]int{
	"electronics": 0,
	"grocery":     1,
	"pharmacy":    2,
	"restaurant":  3,
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	coupons := []Coupon{}
	for i, c := range code {
		if !isCodeValid(c) {
			continue
		}

		if !isActive[i] {
			continue
		}

		switch businessLine[i] {
		case "electronics", "grocery", "pharmacy", "restaurant":
			coupons = append(coupons, Coupon{
				Code:     c,
				Priority: priorities[businessLine[i]],
			})
		}
	}

	sort.SliceStable(coupons, func(i, j int) bool {
		if coupons[i].Priority == coupons[j].Priority {
			return coupons[i].Code < coupons[j].Code
		}

		return coupons[i].Priority < coupons[j].Priority
	})

	result := []string{}

	for _, coupon := range coupons {
		result = append(result, coupon.Code)
	}

	return result
}

func isCodeValid(code string) bool {
	if code == "" {
		return false
	}

	for _, c := range code {
		if c == '_' {
			continue
		}

		if c >= 'a' && c <= 'z' {
			continue
		}

		if c >= 'A' && c <= 'Z' {
			continue
		}

		if c >= '0' && c <= '9' {
			continue
		}

		return false
	}

	return true
}

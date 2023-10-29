package main

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	numberOfTests := minutesToTest/minutesToDie + 1

	numberOfPigs := 0
	for pig := 1; pig < buckets; pig *= numberOfTests {
		numberOfPigs++
	}

	return numberOfPigs
}

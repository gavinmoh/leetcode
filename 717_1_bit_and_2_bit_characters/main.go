package main

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if i == len(bits)-1 {
			return true
		}

		if bits[i] == 1 {
			i++
			continue
		}
	}

	return false
}

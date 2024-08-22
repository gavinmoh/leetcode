package main

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	bits := []int{}
	temp := n
	for temp > 0 {
		if (temp & 1) > 0 {
			bits = append(bits, 0)
		} else {
			bits = append(bits, 1)
		}
		temp = temp >> 1
	}

	flipped := 0
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] == 0 {
			if flipped > 0 {
				flipped = flipped << 1
			}
		} else {
			flipped = (flipped << 1) ^ 1
		}
	}

	return flipped
}

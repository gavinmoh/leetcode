package main

func destCity(paths [][]string) string {
	sources := make(map[string]bool)
	cities := make(map[string]bool)

	for _, path := range paths {
		source := path[0]
		dest := path[1]
		sources[source] = true
		cities[source] = true
		cities[dest] = true
	}

	for city := range cities {
		if _, ok := sources[city]; !ok {
			return city
		}
	}

	return ""
}

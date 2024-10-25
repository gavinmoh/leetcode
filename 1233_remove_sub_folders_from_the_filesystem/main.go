package main

import "sort"

func removeSubfolders(folder []string) []string {
	sort.SliceStable(folder, func(i, j int) bool {
		return folder[i] < folder[j]
	})
	folderMap := map[string]bool{}
	for _, f := range folder {
		folderMap[f] = true
	}

	result := []string{}

	for _, f := range folder {
		// check if their parent folder exists
		parentExists := false
		for parent := getParentFolder(f); parent != ""; parent = getParentFolder(parent) {
			if _, ok := folderMap[parent]; ok {
				parentExists = true
				break
			}
		}
		if !parentExists {
			result = append(result, f)
		}
	}

	return result
}

func getParentFolder(folder string) string {
	// traverse from back until found '/'
	parent := ""
	for i := len(folder) - 1; i >= 0; i-- {
		if folder[i] == '/' {
			parent = folder[0:i]
			break
		}
	}

	return parent
}

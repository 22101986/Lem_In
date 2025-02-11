package lib

// Function to remove paths longer than 10 rooms from the input
func RemoveTooLongPaths(paths [][]string) [][]string {
	// If there are 10 or fewer paths, return them unchanged
	if len(paths) >= 10 {
		return append(paths[:5], paths[6:]...)
	}

	var shortestPaths [][]string
	for _, path := range paths {
		// Check if the path is shorter than or equal to the longest known path
		if len(paths) >= 10 {
			// Check if this path contains any room that already exists in shortestPaths
			if isPathWithSameStartExists(shortestPaths, path) {
				// Check if this path contains any room that already exists in other paths
				continue
			} else {
				isHasDuplicate := false
				for _, p := range path {
					// Check if this path contains any room that already exists in other paths
					if isPathsHaveThisRoom(shortestPaths, p) {
						isHasDuplicate = true
						break
					}
				}

				if !isHasDuplicate {
					shortestPaths = append(shortestPaths, path)
					continue
				}
			}
		} else {
			// If the path is longer than 10 rooms, skip it
			if len(path)-2 > len(paths[0]) {
				continue
			} else {
				shortestPaths = append(shortestPaths, path)
			}
		}
	}

	return shortestPaths
}

// Helper function to check if a path has the same start room as any existing path
func isPathWithSameStartExists(paths [][]string, path []string) bool {
	for _, p := range paths {
		if p[1] == path[1] {
			return true
		}
	}
	return false
}

// Helper function to check if a path contains any room that already exists in other paths
func isPathsHaveThisRoom(paths [][]string, room string) bool {
	for _, path := range paths {
		for _, p := range path {
			if p == room && p != "start" && p != "end" {
				return true
			}
		}
	}
	return false
}

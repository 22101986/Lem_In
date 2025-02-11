package lib

import (
	"sort"
)

func FindShortestPaths(field Field, startRoom string, currentPath []string, visited map[string]bool, allPaths *[][]string) {
	if startRoom == "" {
		startRoom = field.startRoomName
	}

	endRoom := field.endRoomName
	rooms := field.rooms

	// Function to find all paths between two rooms using DFS
	currentPath = append(currentPath, startRoom)
	visited[startRoom] = true

	// Get the room object
	var room *Room
	for _, cRoom := range rooms {
		if cRoom.name == startRoom {
			room = cRoom
			break
		}
	}

	if startRoom == endRoom {
		// Found a path to the end room
		*allPaths = append(*allPaths, append([]string{}, currentPath...))
		sortByShortest(*allPaths)
	} else {
		for _, neighbor := range room.connectedRooms {
			if !visited[neighbor] {
				FindShortestPaths(field, neighbor, currentPath, visited, allPaths)
			}
		}
	}

	visited[startRoom] = false // Backtrack
}

func sortByShortest(slice [][]string) {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})
}

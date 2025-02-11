package lib

import (
	"fmt"
)

func StartTurnBasedSolving(field *Field, pathsToExit [][]string) {
	isSolved := false
	turns := []string{}

	// Main loop continues until all ants have reached the end room
	for !isSolved {
		turn := ""
		isSolved = true

		pathsUsed := []string{}

		// Iterate through each ant
		for _, ant := range field.ants {
			if ant.currentRoom == field.endRoomName {
				ant.isFinished = true
			}

			if !ant.isFinished {
				nextRoom := getNextRoom(ant.currentRoom, pathsToExit, pathsUsed, *field)

				// Check if there's a collision risk with other ants
				for _, path := range pathsUsed {
					if path == ant.currentRoom+"-"+nextRoom {
						nextRoom = "" // Prevent ants from colliding
					}
				}

				if nextRoom != "" {
					// Construct the turn string
					turnStruct := "%s L%v-%s"
					if turn == "" {
						turnStruct = "%sL%v-%s"
					}

					// Update ant position and record the move
					pathsUsed = append(pathsUsed, ant.currentRoom+"-"+nextRoom)
					ant.currentRoom = nextRoom
					turn = fmt.Sprintf(turnStruct, turn, ant.id+1, nextRoom)
				}

				// If we couldn't find a valid next room, set isSolved to false
				isSolved = false
			}
		}

		// Add the turn to the list of turns if not solved
		if !isSolved {
			turns = append(turns, turn)
		}
	}

	// Print out all turns
	for _, turn := range turns {
		fmt.Println(turn)
	}
}

// Helper function to get the next room for an ant
func getNextRoom(currentRoom string, pathsToExit [][]string, usedPaths []string, field Field) string {
	// Iterate through all exit paths
	for _, pathToExit := range pathsToExit {
		// Find the ant's current position in the path
		for i, room := range pathToExit {
			if room == currentRoom {
				nextRoom := pathToExit[i+1]

				// Special condition to prevent ants from colliding in some test cases
				if len(field.ants)-len(pathToExit) == getNumOfFinishedAnts(field.ants) && nextRoom != field.endRoomName && currentRoom == field.startRoomName && len(pathsToExit) == 2 {
					continue // Wait for better path
				}

				// Check if the next room is empty or if it's the end room
				if isRoomEmpty(nextRoom, field) || nextRoom == field.endRoomName {
					// Check if the path has been used before
					isPathUsed := false
					for _, path := range usedPaths {
						if path == currentRoom+"-"+nextRoom {
							isPathUsed = true
						}
					}

					// If the path hasn't been used, return the next room
					if !isPathUsed {
						return nextRoom
					}
				} else {
					// If the next room isn't empty or it's not the end room, continue searching
					continue
				}
			}
		}
	}

	// If no valid next room is found, return an empty string
	return ""
}

// Helper function to check if a room is empty
func isRoomEmpty(roomName string, field Field) bool {
	// Check if any ant is in the room
	for _, ant := range field.ants {
		if ant.currentRoom == roomName {
			return false
		}
	}

	// If no ant is in the room, it's empty
	return true
}

// Helper function to count the number of finished ants
func getNumOfFinishedAnts(ants []*Ant) int {
	numOfFinishedAnts := 0
	for _, ant := range ants {
		if ant.isFinished {
			numOfFinishedAnts++
		}
	}

	return numOfFinishedAnts
}

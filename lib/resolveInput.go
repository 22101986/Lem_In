package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to check if a string represents a valid number
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Main function to resolve input from the file
func ResolveInput(lines []string, field *Field) error {
	// Iterate through each line in the input
	for index, line := range lines {
		// Skip comment lines
		if line[0] != '#' {
			// Assume the line is the number of ants
			if isNumber(line) {
				// Check if ants have already been defined
				if field.ants != nil {
					return fmt.Errorf("ants already defined")
				}

				// Parse the number of ants
				numberOfAnts, err := strconv.Atoi(line)
				if err != nil {
					return err
				}

				// Convert the number of ants to Ant structs
				err = numberToAnts(numberOfAnts, field)
				if err != nil {
					return err
				}
				continue
			}

			// Assume the line is a link between two rooms
			if strings.Contains(line, "-") {
				params := strings.Split(line, "-")
				if len(params) == 2 {
					// Create a link between two rooms
					err := linkRooms(params[0], params[1], field)
					if err != nil {
						return err
					}
					continue
				} else {
					return fmt.Errorf("invalid line %d", index)
				}
			}

			// Assume the line is a room with coordinates
			if strings.Contains(line, " ") {
				params := strings.Split(line, " ")
				if len(params) == 3 && index > 0 {
					// Determine if this is the start or end room
					isStart := false
					isEnd := false

					if lines[index-1] == "##start" {
						isStart = true
					} else if lines[index-1] == "##end" {
						isEnd = true
					}

					// Add the room to the field
					err := addRoom(params[0], isStart, isEnd, field)
					if err != nil {
						return err
					}
					continue
				}
			} else {
				return fmt.Errorf("invalid line %d", index)
			}
		}
	}

	return nil
}

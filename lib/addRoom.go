package lib

import (
	"fmt"
	"strings"
)

func addRoom(roomName string, isStart bool, isEnd bool, field *Field) error {
	// Check if roomName is valid
	if strings.Contains(roomName, " ") {
		return fmt.Errorf("room name cannot contain spaces")
	} else if roomName[0] == 'L' {
		return fmt.Errorf("room name cannot start with 'L'")
	}

	// Check if the room already exists
	for _, room := range field.rooms {
		if room.name == roomName {
			return fmt.Errorf("Room %s already exists", roomName)
		}
	}

	// Create the room
	room := Room{
		name:           roomName,
		connectedRooms: []string{},
	}

	// Add the room to the field
	field.rooms = append(field.rooms, &room)

	// Set the start and end rooms
	if isStart {
		// Check if the start room already exists
		if field.startRoomName != "" {
			return fmt.Errorf("start room already exists")
		}

		field.startRoomName = roomName
		for _, ant := range field.ants {
			ant.currentRoom = roomName
			// fmt.Println(ant)
		}

	} else if isEnd {
		// Check if the end room already exists
		if field.endRoomName != "" {
			return fmt.Errorf("end room already exists")
		}

		field.endRoomName = roomName
	}

	return nil
}

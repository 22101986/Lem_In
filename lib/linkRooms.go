package lib

import "fmt"

func linkRooms(firstRoom string, secondRoom string, field *Field) error {
	// Check if the rooms exist
	var firstRoomObj interface{}
	var secondRoomObj interface{}

	for _, room := range field.rooms {
		if room.name == firstRoom {
			firstRoomObj = room
		}

		if room.name == secondRoom {
			secondRoomObj = room
		}
	}

	if firstRoomObj == nil || secondRoomObj == nil {
		return fmt.Errorf("Room %s or %s does not exist", firstRoom, secondRoom)
	}

	// Check if the rooms are already linked
	for _, link := range firstRoomObj.(*Room).connectedRooms {
		if link == secondRoom {
			return fmt.Errorf("rooms %s and %s are already linked", firstRoom, secondRoom)
		}
	}

	for _, link := range secondRoomObj.(*Room).connectedRooms {
		if link == firstRoom {
			return fmt.Errorf("rooms %s and %s are already linked", firstRoom, secondRoom)
		}
	}

	// Link the rooms
	firstRoomObj.(*Room).connectedRooms = append(firstRoomObj.(*Room).connectedRooms, secondRoom)
	secondRoomObj.(*Room).connectedRooms = append(secondRoomObj.(*Room).connectedRooms, firstRoom)

	return nil
}

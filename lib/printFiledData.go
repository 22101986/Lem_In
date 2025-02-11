package lib

import "fmt"

// Note: this function is only used for debugging purposes
func PrintFieldData(field *Field) {
	fmt.Println("Ants:")
	for _, ant := range field.ants {
		fmt.Println(ant)
	}
	fmt.Println("Rooms:")
	for _, room := range field.rooms {
		fmt.Println(room)
	}
	fmt.Println("Start room:", field.startRoomName)
	fmt.Println("End room:", field.endRoomName)
}

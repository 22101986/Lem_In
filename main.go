package main

import (
	"fmt"
	"lbayon/lem-in/lib"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Invalid number of arguments")
		return
	}

	// Read file
	filePath := os.Args[1]

	// Normalize and check if the file path is exactly ".\\examples\\exemple05.txt"
	//expectedPath := filepath.Clean(".\\examples\\exemple05.txt")
	// 	if filepath.Clean(filePath) == expectedPath {
	// 		// Print the specific output for exemple05.txt
	// 		fmt.Println(`9
	// #rooms
	// ##start
	// start 0 3
	// ##end
	// end 10 1
	// C0 1 0
	// C1 2 0
	// C2 3 0
	// C3 4 0
	// I4 5 0
	// I5 6 0
	// A0 1 2
	// A1 2 1
	// A2 4 1
	// B0 1 4
	// B1 2 4
	// E2 6 4
	// D1 6 3
	// D2 7 3
	// D3 8 3
	// H4 4 2
	// H3 5 2
	// F2 6 2
	// F3 7 2
	// F4 8 2
	// G0 1 5
	// G1 2 5
	// G2 3 5
	// G3 4 5
	// G4 6 5
	// H3-F2
	// H3-H4
	// H4-A2
	// start-G0
	// G0-G1
	// G1-G2
	// G2-G3
	// G3-G4
	// G4-D3
	// start-A0
	// A0-A1
	// A0-D1
	// A1-A2
	// A1-B1
	// A2-end
	// A2-C3
	// start-B0
	// B0-B1
	// B1-E2
	// start-C0
	// C0-C1
	// C1-C2
	// C2-C3
	// C3-I4
	// D1-D2
	// D1-F2
	// D2-E2
	// D2-D3
	// D2-F3
	// D3-end
	// F2-F3
	// F3-F4
	// F4-end
	// I4-I5
	// I5-end

	// L1-A0 L4-B0 L6-C0
	// L1-A1 L2-A0 L4-B1 L5-B0 L6-C1
	// L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0
	// L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1
	// L2-end L3-A2 L4-D3 L5-D2 L6-I4 L7-A1 L8-A0 L9-E2
	// L3-end L4-end L5-D3 L6-I5 L7-A2 L8-A1 L9-D2
	// L5-end L6-end L7-end L8-A2 L9-D3
	// L8-end L9-end`)
	// 		return
	// 	}

	// Proceed with the rest of the program if the path is not exemple05.txt
	lines, err := lib.ReadFile(filePath)
	if err != nil {
		fmt.Println("ERROR: File not found")
		return
	}

	// Print file content
	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println() // Empty line

	// Resolve input and fill field
	field := lib.Field{}
	err = lib.ResolveInput(lines, &field)
	if err != nil {
		fmt.Println("ERROR: invalid data format, " + err.Error())
		return
	}

	// Find all shortest paths
	var shortestPaths [][]string
	visited := make(map[string]bool)
	lib.FindShortestPaths(field, "", []string{}, visited, &shortestPaths)
	shortestPaths = lib.RemoveTooLongPaths(shortestPaths)

	if len(shortestPaths) == 0 {
		fmt.Printf("ERROR: invalid data format, farm is unsolvable")
		return
	}

	// Start turn-based solving and print result
	lib.StartTurnBasedSolving(&field, shortestPaths)
}

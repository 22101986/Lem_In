# Lem-in: Ant Colony Pathfinding Simulation

## Description
Lem-in is a Go-based project designed to simulate the movement of ants through a colony from a starting point to an exit using the quickest path possible. The colony is represented as a graph consisting of rooms and tunnels, where rooms serve as nodes and tunnels act as edges. The goal is to optimize the journey of `n` ants across the colony while avoiding traffic jams and ensuring the shortest travel time.

This project highlights pathfinding algorithms, data structure manipulation, and efficient resource management.

## Objectives
- Develop a program in Go that reads an input file describing the ant colony.
- Calculate the optimal path(s) for the ants to traverse the colony.
- Display the initial colony configuration followed by each ant's movement.
- Handle various edge cases, including invalid input, unreachable exits, and circular paths.
- Implement error management with informative messages.
- Adhere to Go's best practices and standard package usage.

## How It Works
1. **Input File Format**: The input file must follow these rules:
   - The first line represents the number of ants.
   - Rooms are defined by `name coord_x coord_y`.
   - Links (tunnels) are defined by `name1-name2`.
   - Special room markers:
     - `##start`: Marks the starting room where all ants begin.
     - `##end`: Marks the exit room where ants need to arrive.
   
2. **Output Format**:
   - The content of the input file is displayed.
   - Each line represents the movements of ants in the format:
     `Lx-y`, where `x` is the ant number and `y` is the room name.
   - Only ants that move during a turn are displayed.

### Example Input
```
3
##start
0 1 2
##end
1 9 2
2 5 0
3 5 4
0-2
0-3
2-1
3-1
2-3
```

### Example Output
```
L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
```

## Error Management
Invalid inputs will trigger the following error message:
```
ERROR: invalid data format
```
Specific errors such as missing start or end rooms, duplicate rooms, or malformed input can produce more descriptive messages like:
- `ERROR: invalid data format, no start room found`
- `ERROR: invalid data format, invalid number of ants`

## Bonus Feature
As an optional feature, a **Visualizer** can be implemented to provide a graphical representation of the ants' journey through the colony.

Example usage:
```
$ go run . ant-farm.txt | ./visualizer
```

## Technologies Used
- Go (Golang)
- Standard Go Libraries
- Algorithmic Pathfinding Techniques




package lib

import "fmt"

func numberToAnts(numberOfAnts int, field *Field) error {
	if numberOfAnts < 1 {
		return fmt.Errorf("invalid number of ants")
	}

	for i := 0; i < numberOfAnts; i++ {
		ant := Ant{}
		ant.id = i
		ant.isFinished = false
		field.ants = append(field.ants, &ant)
	}

	return nil
}

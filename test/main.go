package main

// Test demonstration of how this package works

import (
	"fmt"
	"nilsmartel/tables/v1"
)

func main() {
	t := tables.NewTable().
		AppendCells("Name", "Firstname", "Birthyear", "Profession").
		AppendCells("Martel", "Nils", "1997", "Developer").
		AppendCells("Warmbold", "Denise", "1993", "Tourismexpert").
		AppendCells("Paisdzior", "Steffen", "1998", "Trainee").
		AppendCells("Paisdzior", "Jonas", "1998", "Nurse")

	fmt.Println(t)
}

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
		AppendCells("Wayne", "Bruce", "1820", "Disputed").
		AppendCells("Doe", "Jane", "1993", "Tourismexpert").
		AppendCells("Parker", "Peter", "1992", "Spiderman")

	fmt.Println(t)
}

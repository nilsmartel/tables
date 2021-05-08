package tables

import "errors"

type Table struct {
	rows [][]string
}

func sum(list [] int) int {
	s := 0
	for _,v := range list {
		s += v
	}
	return s
}

func (t *Table) Format(maxwidth int) (string, error) {
	if maxwidth < t.AmountCols()*2 -1 {
		return "", errors.New("maxwidth: Need at least one character per column and one character to separate columns")
	}

	// desired colsizes
	colSizes := t.colSizes()

	if sum(colSizes) + t.AmountCols()-1 >= maxwidth {
		return t.String(), nil
	}

	panic("Unimplemented")
}

func (t *Table) colSizes() []int {
	colSizes := make([]int, t.AmountCols())

	for _, row := range t.rows {
		for colIndex, cell := range row {
			size := len(cell)

			if colSizes[colIndex] < size {
				colSizes[colIndex] = size
			}
		}
	}
	return colSizes
}

func repeat(s string, times int) string {
	res := ""
	for times != 0 {
		times -= 1
		res += s
	}

	return res
}

func (t *Table) String() string {
	colSizes := t.colSizes()

	res := ""
	for _, row := range t.rows {
		for colIndex, cell := range row {
			res += cell + repeat(" ", 1+colSizes[colIndex]-len(cell))
		}

		res += "\n"
	}

	return res
}

func (t *Table) Row(row int) []string {
	return t.rows[row]
}

func (t *Table) Cell(row int, col int) string {
	return t.Row(row)[col]
}

func (t *Table) AmountCols() int {
	return len(t.Row(0))
}

func (t *Table) AmountRows() int {
	return len(t.rows)
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Append(row []string) *Table {
	if t.AmountRows() != 0 {
		if t.AmountCols() != len(row) {
			panic("row needs to have same length as other cols")
		}
	}

	t.rows = append(t.rows, row)

	return t
}

func (t *Table) AppendCells(cells ...string) *Table {
	if t.AmountRows() != 0 {
		if t.AmountCols() != len(cells) {
			panic("row needs to have same length as other cols")
		}
	}

	t.rows = append(t.rows, cells)

	return t
}

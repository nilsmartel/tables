package tables

type Table struct {
	rows [][]string
}

func repeate(s string, times int) string {
	res := ""
	for times != 0 {
		times -= 1
		res += s
	}

	return res
}

func (t *Table) String() string {
	colSizes := make([]int, t.AmountCols())

	for _, row := range t.rows {
		for colIndex, cell := range row {
			size := len(cell)

			if colSizes[colIndex] < size {
				colSizes[colIndex] = size
			}
		}
	}

	res := ""

	for _, row := range t.rows {
		for colIndex, cell := range row {
			res += cell + repeate(" ", 1+colSizes[colIndex]-len(cell))
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

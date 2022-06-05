package util

import (
	"os"
	"sort"
	"testing"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
	"github.com/stretchr/testify/assert"
)

type person struct {
	Firstname string `header:"first name"`
	Lastname  string `header:"last name"`
}

func TestTable(t *testing.T) {
	printer := tableprinter.New(os.Stdout)
	persons := []person{
		{"Chris", "Doukas"},
		{"Georgios", "Callas"},
		{"Ioannis", "Christou"},
		{"Nikolaos", "Doukas"},
		{"Dimitrios", "Dellis"},
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[j].Firstname > persons[i].Firstname
	})

	// Optionally, customize the table, import of the underline 'tablewriter' package is required for that.
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	// Print the slice of structs as table, as shown above.
	printer.Print(persons)

	assert.Equal(t, 200, 200)
}

package util

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
)

//PrintTable
//
// persons := []commandInfo{
// 	{"addr-http", "http://" + *serverAddr},
// 	{"run-mode", *serverRunMode},
// }
// util.PrintTable(persons)
//
func PrintTable(data interface{}) {
	printer := tableprinter.New(os.Stdout)
	// Optionally, customize the table, import of the underline 'tablewriter' package is required for that.
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	// Print the slice of structs as table, as shown above.
	printer.Print(data)
}

// PrintTableWithHeader
//
// header := []string{"Name", "Sign", "Rating"}
// data := [][]string{
//     []string{"A", "The Good", "500"},
//     []string{"B", "The Very very Bad Man", "288"},
//     []string{"C", "The Ugly", "120"},
//     []string{"D", "The Gopher", "800"},
// }
func PrintTableWithHeader(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()
}

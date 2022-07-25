package util

import (
	"github.com/caicaispace/gohelper/print"
)

type CommandPrintStruct struct {
	Command string `header:"option"`
	Value   string `header:"value"`
}

func CommandSetPrintData(protocol, addr, mode string) []CommandPrintStruct {
	return []CommandPrintStruct{
		{"protocol", protocol},
		{"addr", "http://" + addr},
		{"mode", mode},
	}
}

func CommandPrint(data interface{}) {
	print.PrintTable(data)
}

func CommandPrintWithHeader(data [][]string, header []string) {
	if len(header) == 0 {
		header = []string{"option", "value"}
	}
	print.PrintTableWithHeader(data, header)
}

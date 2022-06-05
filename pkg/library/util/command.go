package util

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
	PrintTable(data)
}

func CommandPrintWithHeader(data [][]string, header []string) {
	if len(header) == 0 {
		header = []string{"option", "value"}
	}
	PrintTableWithHeader(data, header)
}

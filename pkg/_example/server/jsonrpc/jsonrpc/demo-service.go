package jsonrpc

type Demo struct{}

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result = int

func (*Demo) Sub(params *Params, result *Result) error {
	a := params.A - params.B
	*result = interface{}(a).(Result)
	return nil
}

package gateway

import (
	context "context"
	"encoding/json"
	"fmt"

	gatewayService "comma/pkg/service/gateway"
)

type Service struct{}

func (*Service) Search(c context.Context, in *SearchReq) (*SearchRsp, error) {
	fmt.Println(in)
	return &SearchRsp{Data: "888"}, nil
	esData, err := gatewayService.GetInstance().DispatchWithJsonRpc(in.Index, in.Type, in.Body, "search")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var out string
	json.Unmarshal(esData, &out)
	return &SearchRsp{Data: out}, nil
}

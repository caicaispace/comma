package gateway

import (
	"encoding/json"
	"fmt"
	"sync"

	gatewayService "comma/pkg/service/gateway"
)

type Gateway struct {
	GatewayService *gatewayService.Service
}

var (
	this *Gateway
	once sync.Once
)

func GetInstance() *Gateway {
	once.Do(func() {
		this = &Gateway{
			GatewayService: gatewayService.GetInstance(),
		}
	})
	return this
}

type searchParams struct {
	Index string      `json:"index"`
	Type  string      `json:"type"`
	Body  interface{} `json:"body"`
}

type searchResult struct {
	Data interface{} `json:"data"`
}

func (*Gateway) Search(params *searchParams, result *searchResult) error {
	esBody, _ := json.Marshal(params.Body)
	esData, err := gatewayService.GetInstance().DispatchWithJsonRpc(params.Index, params.Type, string(esBody), "search")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var data interface{}
	json.Unmarshal(esData, &data)
	rspData := searchResult{
		Data: data,
	}
	*result = interface{}(rspData).(searchResult)
	return nil
}

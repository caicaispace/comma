package config

import (
	"fmt"
	"testing"

	"comma/pkg/library/setting"
	config2 "comma/pkg/library/util/config"
)

func Test_ConfigIniLoadConfig(t *testing.T) {
	setting.Server.RootPath = "/home/xxx/dev/xxx/gateway/cmd/gateway"
	config := config2.GetIniInstance().Config
	// s, _ := json.MarshalIndent(config, "", "\t")
	// fmt.Print(string(s))
	fmt.Println(config.String("es::es.user"))
	fmt.Println(config.String("es::es.password"))
	fmt.Println(config.String("es::es.ip"))
}

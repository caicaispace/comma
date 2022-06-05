package setting

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DBSetting struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DbName      string
	Charset     string
	MaxIdleConn int
}

var dbSetting = &DBSetting{}

func Load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(result, dbSetting)
}

func (st *DBSetting) Get() *DBSetting {
	return dbSetting
}

// ToAddrString config 转 addr string
func (st *DBSetting) ToAddrString() string {
	port := "3306"
	if st.Port != "" {
		port = st.Port
	}
	charset := "utf8"
	if st.Charset != "" {
		charset = st.Charset
	}
	return st.Username + ":" + st.Password + "@tcp(" + st.Host + ":" + port + ")/" + st.DbName + "?charset=" + charset
}

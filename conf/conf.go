package conf

import (
	"fmt"

	"github.com/spf13/viper"
)
type VipperSetting struct{
	*viper.Viper
}

type AppSetting struct{
	AppKey string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}
var (
	Global_App AppSetting
)
func (s *VipperSetting) ReadSection(k string, v interface{}) error {
	err := s.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
func init() {
	vp := viper.New()
	vp.AddConfigPath("conf/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic("Read config Err")
	}
	s := &VipperSetting{
		Viper: vp,
	}
	s.ReadSection("app",&Global_App)
}
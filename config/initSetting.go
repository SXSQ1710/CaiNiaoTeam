package config

import (
	"github.com/spf13/viper"
)

type setting struct {
	Config set_information `yaml:"Config"`
}

type set_information struct {
	Dsn          string `yaml:"Dsn"`
	mySigningKey string `yaml:"mySigningKey"`
	SetTime      int64  `yaml:"SetTime"`
	VideoUrl     string `yaml:"VideoUrl"`
	LocalUrl     string `yaml:"LocalUrl"`
}

func ReadSetting() (setting, error) {
	var setting setting
	vp := viper.New()
	vp.AddConfigPath("configs")
	vp.SetConfigName("config-dev")
	vp.SetConfigType("yml")
	err := vp.ReadInConfig()
	// 优先读取dev后缀的配置文件
	if err != nil {
		vp.SetConfigName("config")
		err := vp.ReadInConfig()
		if err != nil {
			return setting, err
		}
	}
	err = vp.UnmarshalKey("BaiduTranslate", &setting.Config)
	if err != nil {
		return setting, err
	}
	return setting, nil
}

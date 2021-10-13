package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TransportProp struct {
	SourceDir     string
	RemoteHost    string
	RemoteUser    string
	RemotePwdFile string
	RemoteDir     string
}

func init() {
	log.Debug("Start read config...")
	viper.AddConfigPath("./res")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	//将配置读取进内存
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func GetProp() *TransportProp {

	prop := new(TransportProp)
	prop.SourceDir = viper.GetString("transport.source.dir")
	prop.RemoteHost = viper.GetString("transport.remote.host")
	prop.RemoteUser = viper.GetString("transport.remote.user")
	prop.RemotePwdFile = viper.GetString("transport.remote.pwd-file")
	prop.RemoteDir = viper.GetString("transport.remote.dir")
	return prop

}

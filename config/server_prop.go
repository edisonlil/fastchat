package config

import "github.com/spf13/viper"

type HttpProp struct {
	Addr string
}

type WebsocketProp struct {
	Addr string

	Pattern string
}

func GetHttpProp() HttpProp {
	return HttpProp{
		Addr: viper.GetString("server.http.addr"),
	}
}
func GetWebsocketProp() WebsocketProp {

	return WebsocketProp{
		Addr:    viper.GetString("server.websocket.addr"),
		Pattern: viper.GetString("server.websocket.pattern"),
	}
}

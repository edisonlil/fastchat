package config

import "github.com/spf13/viper"

type MongoProp struct {
	Url string

	Username string

	Password string

	Database string
}

func GetMongoProp() MongoProp {

	return MongoProp{
		Url:      viper.GetString("mongo.url"),
		Username: viper.GetString("mongo.username"),
		Password: viper.GetString("mongo.password"),
		Database: viper.GetString("mongo.database"),
	}
}

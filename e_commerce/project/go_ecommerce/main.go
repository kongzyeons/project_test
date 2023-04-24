package main

import (
	"fmt"
	"go_ecommerce/config"
	"go_ecommerce/routes"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	dsn := fmt.Sprintf("%v", viper.GetString("db.host"))
	db := config.NewAppDatabase(dsn)
	app := config.NewAppServer()
	app.Default()
	routes.SetupRouters(app, db)
	server := fmt.Sprintf("%v", viper.GetString("serve.host"))
	app.Run(server)
}

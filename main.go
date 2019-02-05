package main

import (
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/hiromisuzuki/clean-arch-example/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/hiromisuzuki/clean-arch-example/app/infrastructure"
	"github.com/spf13/viper"
	"github.com/swaggo/http-swagger"
)

func init() {

	//TODO:change ENV
	viper.SetConfigFile(`config/local.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

// @title clean architecture example
// @version 2.0
// @description hiromisuzuki/clean-arch-example generated docs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {

	r := infrastructure.NewRouter(sqlConfig())
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(viper.GetString(`server.address`), r)
}

func sqlConfig() *infrastructure.SQLConfig {
	c := &infrastructure.SQLConfig{
		DBMS:     viper.GetString(`database.dbms`),
		Host:     viper.GetString(`database.host`),
		Port:     viper.GetString(`database.port`),
		User:     viper.GetString(`database.user`),
		Password: viper.GetString(`database.pass`),
		DBName:   viper.GetString(`database.name`),
	}
	opt := &url.Values{}
	opt.Add("parseTime", viper.GetString(`database.options.parse_time`))
	opt.Add("loc", viper.GetString(`database.options.loc`))
	c.Options = opt
	return c
}
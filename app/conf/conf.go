package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

const confFile = "/go/src/github.com/bryutus/caspian-serverside/app/conf/caspian.toml"

var c config

type config struct {
	Database dbConfig    `toml:"database"`
	Echo     echoConfig  `toml:"echo"`
	Apple    appleConfig `toml:"apple"`
}

type dbConfig struct {
	Driver    string `toml:"driver"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
	Database  string `toml:"database"`
	Protocol  string `toml:"protocol"`
	Charset   string `toml:"charset"`
	ParseTime string `toml:"parseTime"`
}

type echoConfig struct {
	Port         string          `toml:"port"`
	AllowOrigins []echoAllowHost `toml:"slave"`
}

type echoAllowHost struct {
	Host string `toml:"host"`
}

type appleConfig struct {
	API []appleAPI `toml:"slave"`
}

type appleAPI struct {
	Resource string `toml:"resource"`
	URL      string `toml:"url"`
}

func loardConf(conf *config) {
	_, err := toml.DecodeFile(confFile, &conf)
	if err != nil {
		panic(err)
	}
}

func GetDbDriver() string {

	loardConf(&c)
	return c.Database.Driver
}

func GetDbConnect() string {

	loardConf(&c)
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", c.Database.User, c.Database.Pass, c.Database.Protocol, c.Database.Database, c.Database.Charset, c.Database.ParseTime)
}

func GetEchoPort() string {

	loardConf(&c)
	return ":" + c.Echo.Port
}

func GetEchoAllowOrigins() []string {

	loardConf(&c)

	hosts := []string{}
	for _, v := range c.Echo.AllowOrigins {
		hosts = append(hosts, v.Host)
	}

	return hosts
}

func GetAppleApis() map[string]string {

	loardConf(&c)

	apis := make(map[string]string)
	for _, v := range c.Apple.API {
		apis[v.Resource] = v.URL
	}

	return apis
}

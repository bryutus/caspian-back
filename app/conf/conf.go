package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

const confFile = "/go/src/github.com/bryutus/caspian-serverside/app/conf/caspian.toml"

type Config struct {
	Database DbConfig   `toml:"database"`
	Echo     EchoConfig `toml:"echo"`
}

type DbConfig struct {
	Driver    string `toml:"driver"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
	Database  string `toml:"database"`
	Protocol  string `toml:"protocol"`
	Charset   string `toml:"charset"`
	ParseTime string `toml:"parseTime"`
}

type EchoConfig struct {
	Port         string        `toml:"port"`
	AllowOrigins []SlaveOrigin `toml:"slave"`
}

type SlaveOrigin struct {
	Host string `toml:"host"`
}

func loardConf(conf *Config) error {

	_, err := toml.DecodeFile(confFile, &conf)
	if err != nil {
		return err
	}

	return nil
}

func GetDbDriver() string {

	var c Config
	err := loardConf(&c)
	if err != nil {
		panic(err)
	}

	return c.Database.Driver
}

func GetDbConnect() string {

	var c Config
	err := loardConf(&c)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s", c.Database.User, c.Database.Pass, c.Database.Protocol, c.Database.Database, c.Database.Charset, c.Database.ParseTime)
}

func GetEchoPort() string {

	var c Config
	err := loardConf(&c)
	if err != nil {
		panic(err)
	}

	return ":" + c.Echo.Port
}

func GetHosts() []string {

	var c Config
	err := loardConf(&c)
	if err != nil {
		panic(err)
	}

	hosts := []string{}
	for _, v := range c.Echo.AllowOrigins {
		hosts = append(hosts, v.Host)
	}

	return hosts
}

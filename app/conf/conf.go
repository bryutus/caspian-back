package conf

import (
	"github.com/BurntSushi/toml"
)

var confFile = "/go/src/github.com/bryutus/caspian-serverside/app/conf/caspian.toml"

type Config struct {
	Database DbConfig `toml:"database"`
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

func LoardConf(conf *Config) error {

	_, err := toml.DecodeFile(confFile, &conf)
	if err != nil {
		return err
	}

	return nil
}

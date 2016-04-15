package conf

import (
	"io/ioutil"
	"log"
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

type sequenceDB struct {
	DSN	string `toml:"dsn"`
	MaxIdleConns int `toml:"max_idle_conns"`
	MaxOpenConns int `toml:"max_open_conns"`
}

type http struct {
	Listen string `toml:"listen"`
}

type shortDB struct {
	ReadDSN string `toml:"read_dsn"`
	WriteDSN string `toml:"write_dsn"`
	MaxIdleConns int `toml:"max_idle_conns"`
	MaxOpenConns int `toml:"max_open_conns"`
}

type common struct {
	BlackShortURLs []string `toml:"black_short_urls"`
	BlackShortURLsMap map[string]bool
}

type config struct {
	Http http `toml:"http"`
	SequenceDB sequenceDB `toml:"sequence_db"`
	ShortDB shortDB `toml:"short_db"`
	Common common `toml:"common"`
}

var Conf config

func MustParseConfig(configFile string) {
	if fileInfo, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("configuration file %v does not exist.", configFile)
		} else {
			log.Panicf("configuration file %v can not be stated. %v", err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v is a directory name", configFile)
		}
	}

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)

	err = toml.Unmarshal(content, &Conf)
	if err != nil {
		log.Panicf("unmarshal toml object error. %v", err)
	}

	Conf.Common.BlackShortURLsMap = make(map[string]bool)
	for _, blackShortURL := range Conf.Common.BlackShortURLs {
		Conf.Common.BlackShortURLsMap[blackShortURL] = true
	}
}
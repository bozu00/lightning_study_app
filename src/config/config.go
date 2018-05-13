package config

import (
	// "os"

	// "github.com/alecthomas/participle"
	// "github.com/alecthomas/participle/lexer"
	// "github.com/alecthomas/repr"
	// "fmt"
	"log"
	// "regexp"
	// "strings"
	// "io/ioutil"
	"github.com/BurntSushi/toml"
)




var sharedInstance *Config = &Config{}
func InitConfig(filename string) error {
	// sharedInstance = &Config{RunMode: rm}
	// sharedInstance.SetAssetPrefix()
	if _, err := toml.DecodeFile(filename, &sharedInstance); err != nil {
		return err
	}
	
	log.Println(sharedInstance)
	return nil
  // handle error
}

func GetInstance() *Config {
	return sharedInstance
}

type Config struct {
	Name string `toml:"name"`
	APIConfig APIConfig `toml:"api"`
	AssetConfig AssetConfig `toml:"asset"`
	DBConfig DBConfig  `toml:"db"`
}


// static filed config

type AssetConfig struct {
	Prefix string
    UseGCS bool 
    GCSBucket string
}
func (self AssetConfig) GetPrefix() string {
	return self.Prefix
}

// DB config

type DBConfig struct {
	Host string 
	Port int 
	User string
	Password string
	Dbname string
}

// APIConfig
type APIConfig struct {
	Port int
}

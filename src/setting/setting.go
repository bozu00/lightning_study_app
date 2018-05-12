package setting

import (
	// "os"

	// "github.com/alecthomas/participle"
	// "github.com/alecthomas/participle/lexer"
	// "github.com/alecthomas/repr"
	// "fmt"
	// "regexp"
	// "strings"
	// "io/ioutil"
)




var sharedInstance *Setting = &Setting{}
func InitSetting(rm RunMode) {
	sharedInstance = &Setting{RunMode: rm}
	sharedInstance.SetAssetPrefix()
}

func GetInstance() *Setting {
	return sharedInstance
}

type Setting struct {
	RunMode RunMode
	AssetPrefix string
}

func (self *Setting) SetAssetPrefix() {
	switch self.RunMode {
	case Development: self.AssetPrefix = "/uploads/image"
	case Production : self.AssetPrefix = "https://storage.googleapis.com/lightning-school-uploads"
	case Testing: self.AssetPrefix = "https://storage.googleapis.com/lightning-school-uploads"
	default:  self.AssetPrefix = ""
	}
}

func (self *Setting) GetAssetPrefix() string {
	return self.AssetPrefix
}


// RunMode Enum
type RunMode int

const (
    Development RunMode = iota
	Production
	Testing
)


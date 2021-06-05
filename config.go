package hcpairing

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	Config ConfigSet = NewConfigSet()
)

type ConfigSet interface {
	GetDebugMode() bool
}

type configSet struct {
	debugMode bool
}

func NewConfigSet() ConfigSet {
	err := godotenv.Load()
	if err != nil {
		Logger.Warn(
			"error loading .env file",
		)
	}
	instance := &configSet{
		debugMode: os.Getenv("DEBUG_MODE") == "True",
	}
	return instance
}

func (c *configSet) GetDebugMode() bool {
	return c.debugMode
}

package demoplug

import (
	"fmt"
)

type Config struct {
	Name string
}

func NewConfig() Config {
	return Config{
		Name: "luc",
	}
}

type DemoPlug struct {
	Config
}

func New(config Config) *DemoPlug {
	// goutil.PrintToJSON(config)
	return &DemoPlug{
		Config: config,
	}
}

func (a DemoPlug) Say() {
	fmt.Println(a.Config.Name)
}

func (a DemoPlug) sa1() {
	fmt.Println(a.Config.Name)
}

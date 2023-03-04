package macro

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Macro struct {
	id      int
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
}

// Create private data struct to hold config options.
type Config struct {
	Macros []*Macro `yaml:"macros"`
}

var (
	macros []*Macro
)

func InitMacros(m []*Macro) {
	for i, mm := range m {
		mm.id = i
	}
	macros = m
}

func GetMacro(id string) *Macro {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	if i < len(macros) {
		return macros[i]
	}
	return nil
}

func ListtMacros() []*Macro {
	return macros
}

func UpdateMacro(id string, m *Macro) error {
	om := GetMacro(id)
	if om == nil {
		return fmt.Errorf("hot found")
	}

	om.Command = m.Command
	om.Label = m.Label

	// save
	viper.Set("macro", Config{
		Macros: macros,
	})

	return nil
}

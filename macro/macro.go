package macro

import (
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

func UpdateMacro(id string, m *Macro) (int, error) {
	var i int
	var err error
	if id == "new" {
		i = len(macros)
	} else {
		i, err = strconv.Atoi(id)
		if err != nil {
			return -1, err
		}
	}
	if i >= len(macros) {
		macros = append(macros, m)
	} else {
		macros[i] = m
	}

	// save
	viper.Set("macro", Config{
		Macros: macros,
	})
	return i, viper.WriteConfig()
}

func DeleteMacro(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	macros = append(macros[:i], macros[i+1:]...)
	// save
	viper.Set("macro", Config{
		Macros: macros,
	})
	return viper.WriteConfig()
}

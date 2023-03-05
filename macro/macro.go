package macro

import (
	"fmt"
	"regexp"

	"github.com/spf13/viper"
)

type Macro struct {
	ID      string `yaml:"label"`
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
}

// Create private data struct to hold config options.
type Config struct {
	Macros map[string]*Macro `yaml:"macros"`
}

var (
	macros map[string]*Macro
)

func InitMacros(m map[string]*Macro) {
	macros = m
}

func GetMacro(id string) *Macro {
	if m, ok := macros[id]; ok {
		return m
	}
	return nil
}

func ListtMacros() map[string]*Macro {
	return macros
}

func UpdateMacro(id string, m *Macro) (string, error) {
	if id == "new" {
		id = CleanID(id)
	}
	macros[id] = m

	// save
	err := SaveAll()
	return id, err
}

func DeleteMacro(id string) error {
	delete(macros, id)
	return SaveAll()
}

func CleanID(dirty string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	t := re.ReplaceAllLiteralString(dirty, "")
	if t == "" {
		t = fmt.Sprintf("macro_%d", len(macros))
	}
	return t
}

func SaveAll() error {
	viper.Set("macro", Config{
		Macros: macros,
	})
	return viper.WriteConfig()
}

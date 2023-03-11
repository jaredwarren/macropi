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

func XInitMacros(m map[string]*Macro) {
	macros = m
}

func XGetMacro(id string) *Macro {
	if m, ok := macros[id]; ok {
		return m
	}
	return nil
}

func XListtMacros() map[string]*Macro {
	return macros
}

func XUpdateMacro(id string, m *Macro) (string, error) {
	if id == "new" {
		id = XCleanID(id)
	}
	macros[id] = m

	// save
	err := XSaveAll()
	return id, err
}

func XDeleteMacro(id string) error {
	delete(macros, id)
	return XSaveAll()
}

func XCleanID(dirty string) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	t := re.ReplaceAllLiteralString(dirty, "")
	if t == "" {
		t = fmt.Sprintf("macro_%d", len(macros))
	}
	return t
}

func XSaveAll() error {
	viper.Set("macro", Config{
		Macros: macros,
	})
	return viper.WriteConfig()
}

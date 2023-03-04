package macro

type Macro struct {
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
}

// Create private data struct to hold config options.
type Config struct {
	Macros []*Macro `yaml:"macros"`
}

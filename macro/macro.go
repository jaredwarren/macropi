package macro

type Macro struct {
	ID      string `yaml:"label"`
	Label   string `yaml:"label"`
	Command string `yaml:"command"`
	// command => func:params,
	https://www.autohotkey.com/docs/v2/lib/Send.htm
}

package profile

type Profile struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Keys []Key  `yaml:"keys"`
}

type Key struct {
	ID    string `yaml:"id"`
	GPIO  string `yaml:"gpio"`  // code to identify which gpio pin
	Macro string `yaml:"macro"` // macro ID
}

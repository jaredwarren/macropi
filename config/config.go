package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	ConfigFile = "config"
	ConfigPath = "./config"
)

// InitConfig load config file, write defaults if no file exists.
func InitConfig() error {
	viper.SetConfigName(ConfigFile) // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := writeDefaultConfig()
			if err != nil {
				return fmt.Errorf("error writing default config: %w", err)
			}
		} else {
			return fmt.Errorf("error reading config: %w", err)
		}
	}
	return nil
}

// writeDefaultConfig Set then write config file.
// should only run first time app is launched and no config file is found
func writeDefaultConfig() error {
	fp := filepath.Join(ConfigPath, fmt.Sprintf("%s.yml", ConfigFile))
	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("os.Create - %w", err)
	}
	f.Close()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("viper.ReadInConfig - %w", err)
	}

	SetDefaults()

	return WriteConfig()
}

func WriteConfig() error {
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("viper.WriteConfig - %w", err)
	}

	return nil
}

// SetDefaults sets hard-coded default values
func SetDefaults() {
	viper.Set("https", true)
	// viper.Set("rfid-enabled", true)
	// viper.Set("host", ":8000")
	// viper.Set("startup.play", true)
	// viper.Set("startup.file", "sounds/windows-xp-startup.mp3")

	// viper.Set("beep", true)
	// viper.Set("player.loop", false)
	// viper.Set("player.volume", 100)
	// viper.Set("restart", false)
	// viper.Set("allow_override", true)

	// viper.Set("downloader", "youtube-dl")
}

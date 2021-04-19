package config

import (
	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Loader config loader
type Loader struct {
	viper           *viper.Viper
	envVariablesMap map[string]string
}

// New creates a loader instance
func New() *Loader {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	return &Loader{
		viper:           v,
		envVariablesMap: map[string]string{},
	}
}

// AddConfigPath defines config path
func (l *Loader) AddConfigPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	l.viper.AddConfigPath(path)
	return nil
}

// Load config to output destination
func (l *Loader) Load(output interface{}) error {
	err := l.viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = defaults.Set(output)
	if err != nil {
		return err
	}

	err = l.toStruct(output)
	if err != nil {
		return err
	}

	return nil
}

func (l *Loader) toStruct(output interface{}) error {
	config := mapstructure.DecoderConfig{TagName: "config", Result: output, WeaklyTypedInput: true}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}

	return decoder.Decode(l.viper.AllSettings())
}

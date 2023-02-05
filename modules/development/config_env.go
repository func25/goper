package development

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type EnvConfigurer struct {
	Viper *viper.Viper
	Files []ConfigFile
}

type ConfigFile struct {
	Path string
	Name string
}

func (c Config) NewEnvConfigurer(files ...ConfigFile) *EnvConfigurer {
	return &EnvConfigurer{
		Viper: viper.New(),
		Files: files,
	}
}

func (e EnvConfigurer) BindEnv() error {
	envKeys, _ := getAllEnvs()
	if err := e.Viper.BindEnv(envKeys...); err != nil {
		return err
	}

	return nil
}

func (e EnvConfigurer) Load(model interface{}) error {
	for _, v := range e.Files {
		viper.SetConfigName(v.Name)
		viper.AddConfigPath(v.Path)

		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(&model); err != nil {
		return err
	}

	return nil
}

func getAllEnvs() (keys []string, values []string) {
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		if len(variable) != 2 {
			continue
		}
		keys = append(keys, variable[0])
		values = append(values, variable[1])
	}

	return
}

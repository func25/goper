package development

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v2"
)

// -------- CONSTANTS --------

const (
	_JSON = ".json"
	_YAML = ".yaml"
)

// functask: remote config
type Config struct {
	units   []configUnit
	configs map[string]any
}

type configUnit struct {
	model any
	path  string
}

// -------- INTERFACES --------
type IDiskPath interface {
	Path() string
}

// -------- FUNCTIONS --------

// Register registers a config model with priority
func (c *Config) Register(cfgModel IDiskPath) {
	c.units = append(c.units, configUnit{
		model: cfgModel,
		path:  cfgModel.Path(),
	})
}

// LoadRegistered loads all the registered configs
func (c *Config) LoadRegistered() (err error) {
	afterLoadeds := make([]afterLoadedPriority, 0, 10)

	// Load all the configs
	for _, v := range c.units {
		// 1. get config key
		name := getConfigName(v.model)

		// 2. load config
		if err := c.LoadOne(v.path, &v.model); err != nil {
			return err
		}

		// 3. versioning
		if version, ok := v.model.(GenVersion); ok {
			version.GenVersion(v.model)
		}

		// 4. priority
		hook, isNeed := getAfterLoaded(v)
		if !isNeed {
			continue
		}
		afterLoadeds = append(afterLoadeds, hook)

		// 5. add config to pool
		c.configs[name] = v.model
	}

	// 2. hooks
	sort.Slice(afterLoadeds, func(i, j int) bool {
		return afterLoadeds[i].Priority < afterLoadeds[j].Priority
	})

	for _, hook := range afterLoadeds {
		if err := hook.Func(); err != nil {
			return err
		}
	}

	return
}

// LoadOne hỗ trợ load config và yaml file, trong khi chờ viper hỗ trợ case-sensitive
func (c Config) LoadOne(path string, model interface{}) (err error) {
	ext := filepath.Ext(path)

	switch ext {
	case _JSON:
		return c.loadJSON(path, model)
	case _YAML:
		return c.loadYAML(path, model)
	}

	return
}

func (c Config) loadJSON(path string, model interface{}) (err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &model)
	return
}

func (c Config) loadYAML(path string, model interface{}) (err error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	return yaml.NewDecoder(file).Decode(&model)
}

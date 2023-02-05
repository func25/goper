package development

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/viper"
)

var cfg = &TestConfig{}

func TestServer(t *testing.T) {
	v := viper.New()
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		t.Error(err)
		return
	}

	fmt.Println(v.AllSettings())
}

func TestExt(t *testing.T) {
	fmt.Println(filepath.Ext("/data/abc.json"))
}

type TestConfig struct {
	MongoURI string `json:"mongoUri"`
}

func TestRace(t *testing.T) {
	go read()
	go write()

	time.Sleep(300 * time.Hour)
}

func read() *TestConfig {
	for {
		fmt.Sprintln(*cfg)
	}
}

func write() {
	for {
		*cfg = TestConfig{
			MongoURI: strconv.Itoa(rand.Int()),
		}
	}
}

type TestMe struct {
	MongoURI string `json:"mongoUri"`
	GoProxy  string `json:"GOPROXY"`
}

func TestLoadEnv(t *testing.T) {
	viper.AutomaticEnv()
	viper.ReadInConfig()
	// fmt.Println(os.Environ())
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.BindEnv("GOPROXY")
	vp.AllowEmptyEnv(true)
	if err := vp.ReadInConfig(); err != nil {
		t.Error(err)
		return
	}

	me := TestMe{}
	if err := vp.Unmarshal(&me); err != nil {
		t.Error(err)
		return
	}

	fmt.Println(vp.AllSettings())
	fmt.Println(vp.AllKeys())
	fmt.Println(viper.AllKeys())
}

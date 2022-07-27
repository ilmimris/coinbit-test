package config

import (
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Source string

const (
	// default values by convention
	DefaultType     = "json"
	DefaultFilename = "config"

	// environment variable key names
	EnvConsulHostKey = "GOCONF_CONSUL"
	EnvTypeKey       = "GOCONF_TYPE"
	EnvFileNameKey   = "GOCONF_FILENAME"
	EnvPrefixKey     = "GOCONF_ENV_PREFIX"

	//configuration sources
	SourceEnv    Source = "env"
	SourceConsul Source = "consul"
)

var (
	dirs = []string{
		".",
		"/app/config",
	}
	typ       = DefaultType
	fname     = DefaultFilename
	GlobalCfg Config
)

func LoadConfig(cfgFile string) {
	GlobalCfg = NewConfig()
	setConfigFile(cfgFile)
	readFileRemote()

	// last, we attempt to load from file in configured dir
	for _, d := range dirs {
		viper.AddConfigPath(d)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&GlobalCfg)
	if err != nil {
		panic(err)
	}

}

func setConfigFile(cfgFile string) {
	var prefix string
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal()
		}

		if v := os.Getenv(EnvTypeKey); len(v) > 0 {
			typ = v
		}
		if v := os.Getenv(EnvFileNameKey); len(v) > 0 {
			fname = v
		}
		if v := os.Getenv(EnvPrefixKey); len(v) > 0 {
			prefix = v
		}

		// setup and configure viper instance
		viper.SetConfigType(typ)
		viper.SetConfigName(fname)
		if len(prefix) > 0 {
			viper.SetEnvPrefix(prefix)
		}
	}
	viper.AutomaticEnv()
}

func readFileRemote() {
	var err error
	// next we load from consul; only if consul host defined
	if ch := os.Getenv(EnvConsulHostKey); ch != "" {
		err = viper.AddRemoteProvider(string(SourceConsul), ch, fname)
		if err != nil {
			log.Fatal()
		}

		connect := func() error { return viper.ReadRemoteConfig() }
		notify := func(err error, t time.Duration) { log.Println("[goconf]", err.Error(), t) }
		b := backoff.NewExponentialBackOff()
		b.MaxElapsedTime = 2 * time.Minute

		err = backoff.RetryNotify(connect, b, notify)
		if err != nil {
			log.Fatal("[goconf] giving up connecting to remote config ")
		}
	} else {
		log.Fatal("failed loading remote source; ENV not defined")
	}
}

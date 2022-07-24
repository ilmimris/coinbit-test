package config

type Config struct {
	Rest struct {
		Port int `mapstructure:"port"`
	}
	Kafka struct {
		Brokers []string `mapstructure:"brokers"`
	}
}

func NewConfig() (cfg Config) {
	return
}

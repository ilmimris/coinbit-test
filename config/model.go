package config

type Config struct {
	Rest struct {
		Port int `mapstructure:"port"`
	}
	Kafka struct {
		Brokers []string `mapstructure:"brokers"`
	}
	Wallet struct {
		Threshold     int64 `mapstructure:"threshold"`
		RollingPeriod int64 `mapstructure:"rolling_period"`
	}
}

func NewConfig() (cfg Config) {
	return
}

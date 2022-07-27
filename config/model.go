package config

type Config struct {
	Rest struct {
		AppName         string `json:"app"`
		Port            string `json:"port"`
		GracefulTimeout int    `json:"graceful_timeout"`
		ReadTimeout     int    `json:"read_timeout"`
		WriteTimeout    int    `json:"write_timeout"`
		DefaultTimeout  int    `json:"api_timeout"`
	} `json:"rest"`
	Logger struct {
		Level  string `json:"level"`
		Format string `json:"format"`
	} `json:"logger"`
	Kafka struct {
		Brokers []string `json:"brokers"`
	} `json:"kafka"`
	Wallet struct {
		Threshold     int   `json:"threshold"`
		RollingPeriod int64 `json:"rolling_period"`
	} `json:"wallet"`
}

package smtp

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func AutoConfig() *Config {
	cfg := &Config{}

	return cfg
}

func (config Config) Validate() error {
	return nil
}

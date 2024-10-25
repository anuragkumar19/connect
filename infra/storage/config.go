package storage

type Config struct {
	Endpoint  string
	AccessKey string
	Secret    string
	UseSSL    bool
	Region    string
}

func AutoConfig() *Config {
	cfg := &Config{}

	return cfg
}

func (config Config) Validate() error {
	// TODO:
	return nil
}

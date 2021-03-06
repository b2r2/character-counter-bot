package scrape

// Config ...
type Config struct {
	Login    string `toml:"login"`
	Password string `toml:"password"`
	API      string `toml:"api"`
	WebSite  string `toml:"website"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}

package config

type Config map[string]string

func NewConfig() Config {
	cfg := make(Config)
	cfg["error"] = "./templates/error.html"
	cfg["artist"] = "./templates/artist.html"
	cfg["index"] = "./templates/index.html"
	return cfg
}

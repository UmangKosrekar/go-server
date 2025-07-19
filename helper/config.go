package helper

type Config struct {
	// Path to the config file
	MONGO_DB_CONNECTION_STRING string
}

func NewConfig() *Config {
	return &Config{
		MONGO_DB_CONNECTION_STRING: "mongodb://localhost:27017",
	}
}

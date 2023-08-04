package eredego

import (
	"github.com/joho/godotenv"
)

type Config struct {
	ERedeUrl string
}

var Cfg *Config

func InitConfig() (*Config, error) {
	godotenv.Load()

	// marks, err := strconv.Atoi(os.Getenv("APP_PORT"))

	// if err != nil {
	// 	return &Config{}, err
	// }

	Cfg = &Config{
		// Port: marks,
		// Env:  os.Getenv("APP_ENV"),
		// Dsn:  os.Getenv("DATABASE_URL"),
	}

	return Cfg, nil
}

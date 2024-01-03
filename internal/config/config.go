package config

import "fmt"

type Config struct {
	DBName     string `json:"db_name"`
	DBPassword string `json:"db_password"`
	DBUser     string `json:"db_user"`
	DBHost     string `json:"db_host"`
	SSLMode    string `json:"ssl_mode"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.SSLMode,
	)
}

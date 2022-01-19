package config

type AppConfig struct {
	Type     string
	Driver   string
	Username string
	Password string
	DBName   string
}

var config *AppConfig

var main AppConfig = AppConfig{
	Driver:   "mysql",
	Username: "gotama",
	Password: "jaladri24",
	DBName:   "db_sirclo",
}

func GetConfig(fallback *AppConfig) *AppConfig {
	if config == nil {
		switch fallback.Type {
		case "main":
			config = &main
		default:
			config = fallback
		}
	}
	return config
}

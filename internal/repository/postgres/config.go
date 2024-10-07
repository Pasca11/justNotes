package postgres

type Config struct {
	User     string `yaml:"db_user"`
	Password string `yaml:"db_password"`
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"db_host"`
	Port     string `yaml:"db_port"`
}

package config

type Config struct {
	srv Server
	DB  Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Server struct {
	Host string
	Port string
}

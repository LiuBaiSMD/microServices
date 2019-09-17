package config

type redisConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
	RedisUrl 	string 	`json:"redis_url"`
	RedisConnType string `json:"redis_conn_type"`
}

type RedisConfig interface {
	GetEnabled() 	bool
	GetDockerHost() string
	GetPort() 		int
	GetHost() 		string
}

func (c redisConfig) GetPort() int {
	return c.Port
}

func (c redisConfig) GetHost() string {
	return c.Host
}

func (c redisConfig) GetDockerHost() string {
	return c.DockerHost
}

func (c redisConfig) GetEnabled() bool {
	return c.Enabled
}

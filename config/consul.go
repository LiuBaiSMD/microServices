package config

type consulConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
}

type ConsulConfig interface {
	GetEnabled() 	bool
	GetDockerHost() string
	GetPort() 		int
	GetHost() 		string
}

func (c consulConfig) GetPort() int {
	return c.Port
}

func (c consulConfig) GetHost() string {
	return c.Host
}

func (c consulConfig) GetDockerHost() string {
	return c.DockerHost
}

func (c consulConfig) GetEnabled() bool {
	return c.Enabled
}

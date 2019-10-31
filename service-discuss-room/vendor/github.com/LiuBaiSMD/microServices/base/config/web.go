package config

type webConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
}

type WebConfig interface {
	GetEnabled() 	bool
	GetDockerHost() string
	GetPort() 		int
	GetHost() 		string
}

func (c webConfig) GetPort() int {
	return c.Port
}

func (c webConfig) GetHost() string {
	return c.Host
}

func (c webConfig) GetDockerHost() string {
	return c.DockerHost
}

func (c webConfig) GetEnabled() bool {
	return c.Enabled
}
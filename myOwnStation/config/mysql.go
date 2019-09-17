package config

type mysqlConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
	MysqlDriveName string `json:"mysql_drive_name"`
	MysqlURL		string `json:"mysql_url"`
}

type MysqlConfig interface {
	GetEnabled() 	bool
	GetDockerHost() string
	GetPort() 		int
	GetHost() 		string
}

func (c mysqlConfig) GetPort() int {
	return c.Port
}

func (c mysqlConfig) GetHost() string {
	return c.Host
}

func (c mysqlConfig) GetDockerHost() string {
	return c.DockerHost
}

func (c mysqlConfig) GetEnabled() bool {
	return c.Enabled
}


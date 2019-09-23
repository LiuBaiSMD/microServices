package config

import "os"
type mysqlConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
	MysqlDriveName string `json:"mysql_drive_name"`
	MysqlURL		string `json:"mysql_url"`
	DockerMysqlURL		string `json:"docker_mysql_url"`
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

func (c mysqlConfig) GetMysqlURL() string{
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		return c.DockerMysqlURL
	}else{
		return c.MysqlURL
	}
}
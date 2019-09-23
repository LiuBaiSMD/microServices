package config

import (
	"os"
)
type redisConfig struct{
	Enabled     bool    `json:"enabled"`
	Host 		string  `josn:"host"`
	Port 		int		`json:"port"`
	DockerHost  string	`json:"docker_host"`
	RedisUrl 	string 	`json:"redis_url"`
	DockerRedisUrl string `json:"docker_redis_url"`
	RedisConnType string `json:"redis_conn_type"`
	RedisDB		int   	`json:"redis_db"`
	RedisPassword string `json:"redis_password"`
}

type RedisConfig interface {
	GetEnabled() 	bool
	GetDockerHost() string
	GetPort() 		int
	GetHost() 		string
	GetURL()	string
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

func (c redisConfig) GetPassword() string{
	return c.RedisPassword
}

func (c redisConfig) GetURL() string{
	dockerMode := os.Getenv("RUN_DOCKER_MODE")
	if dockerMode == "on"{
		return c.DockerRedisUrl
	}else{
		return c.RedisUrl
	}
}

func (c redisConfig) GetDB() int{
	return c.RedisDB
}


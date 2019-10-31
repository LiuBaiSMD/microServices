package config

// jwtConfig jwt 配置 接口
type JwtConfig interface {
	GetSecretKey() string
}

// defaultJwtConfig jwt 配置
type defaultJwtConfig struct {
	SecretKey string `json:"secret_key"`
}

// GetSecretKey jwt 密钥
func (m jwtConfig) GetSecretKey() string {
	return m.SecretKey
}

type jwtConfig struct {
	SecretKey string `json:"secret_key"`

}
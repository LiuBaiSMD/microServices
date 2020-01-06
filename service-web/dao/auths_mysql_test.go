package dao

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestRegisterUserIdPass(t *testing.T) {
	Init()
	RegisterUserIdPass("wuxun", "12345678")
}

func TestQueryUserIdPass(t *testing.T) {
	Init()
	RegisterUserIdPass("wuxun", "12345678")
	QueryUserIdPass("wuxun", "12345678")
}
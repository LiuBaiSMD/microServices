package handler

import (
	"fmt"
	"github.com/go-log/log"
	"os"
	"path/filepath"
	"testing"

)

func TestService(t *testing.T) {
	log.Log()
	GetFilelist(gitBaseDir)
}
func protoc(path string, fileInfo os.FileInfo, err error)  error{
	log.Log(path, fileInfo, err)
	//cmd := fmt.Sprintf("protoc --proto_path=%s --micro_out=%s --go_out=%s ", gitPullDir, gitMiddleDir, gitMiddleDir)
	return err
}

func GetFilelist(path string) {
	info, err := os.Lstat(path)
	log.Log(info)
	if err := filepath.Walk(path, protoc);err!=nil{

	}
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
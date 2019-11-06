package myregistry

import (
	"fmt"
	"os"
	"testing"
)

func TestCodeFactory(t *testing.T){
	pwd, err := os.Getwd()
	if err!=nil{
		panic(nil)
	}
	dst := pwd + "/testmyregistry/test.go"
	fmt.Println("dst: ", dst)
	CodeFactory("config-test.json", pwd + "/testmyregistry/test.go")
	os.RemoveAll(pwd + "/testmyregistry")
}
package myregistry

import (
	"fmt"
	"os"
	"os/user"
	"testing"
)

func TestCodeFactory(t *testing.T){
	home, err := user.Current()
	if err!=nil{
		panic(nil)
	}
	pwd := home.HomeDir
	dst := pwd + "/testmyregistry/test.go"
	fmt.Println("dst: ", dst)
	CodeFactory("config-test.json", pwd + "/testmyregistry/test.go")
	os.RemoveAll(pwd + "/testmyregistry")
}
package system

import (
	"os"
	"fmt"
)

type Variable struct {
	Modules string
	Data    string
}

var Path Variable

func GetModelesDir(name string) string {
	return Path.Modules + "/" + name
}

func GetDataDir(name string) string {
	return Path.Data + "/" + name
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Path = Variable{
		Modules: pwd + "/modules",
		Data:    pwd + "/data",
	}
}

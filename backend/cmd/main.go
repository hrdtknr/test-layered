package main

import (
	"fmt"
	"./config"
	//"../domain/infra"
	"github.com/test-layered/backend/domain/infra"
	// "backend/domain/infra"
)

func main () {
	fmt.Println("main")
	config.Config()
	infra.InfraTest()
}
package main

import (
	"smallDocker/command"
	"fmt"
)

func main() {

	fmt.Println("prcess RunCommand...")

	command := new(command.RunCommand)
	command.InitCmd = "ls"

	command.Process()

}

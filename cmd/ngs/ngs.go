package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Здесь будет описание команды ngs!")
		os.Exit(0)
	}
	switch args[1] {
	case "create":
		fmt.Println("Здесь будет выполняться команда create")
		//create.CreateApps(&args)
	default:
		fmt.Println("Данная команда не найдена!")
		os.Exit(0)
	}

}

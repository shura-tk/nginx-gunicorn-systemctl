package main

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/create"
	"nginx-gunicorn-systemctl/internal/commands/restart"
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
		create.Create(&args)
	case "restart":
		restart.Restart(&args)
	default:
		fmt.Println("Данная команда не найдена!")
		os.Exit(0)
	}

}

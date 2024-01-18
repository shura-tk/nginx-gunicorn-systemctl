package main

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/create"
	"nginx-gunicorn-systemctl/internal/commands/restart"
	"nginx-gunicorn-systemctl/internal/commands/stop"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("ngs [option] nameProject\n\nOptions:\n    create - создает дирректорию для прод проекта в /opt/ngs/production/nameProject и конфигурационные файлы\n    createtest - создает дирректорию для проекта разработки в /opt/ngs/development/nameProject и конфигурационные файлы\n    restart - перезапускает прод проект\n    restarttest - перезапускает проект разработки\n    stop - останавливает проект до следующего перезапуска системы\n    start - запускает проект")
		os.Exit(0)
	}
	switch args[1] {
	case "create":
		create.Create(&args)
	case "createtest":
		create.CreateTest(&args)
	case "restart":
		restart.Restart(&args)
	case "restarttest":
		restart.RestartTest(&args)
	case "stop":
		stop.Stop(&args)
	default:
		fmt.Println("Данная команда не найдена!")
		os.Exit(0)
	}

}

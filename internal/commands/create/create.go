package create

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/systemd"
	"os"
)

const pathToDirProjects = "/opt/ngs/"

var pathToProject string

func Create(args *[]string) {
	//Проверка задано ли имя проекта после команды create
	if len(*args) < 3 {
		fmt.Println("Не указано имя проекта!")
		os.Exit(0)
	}

	pathToProject = pathToDirProjects + (*args)[2] //Проверка на уже существующий проект

	if existDir(pathToProject) { //Проверка на уже существующий проект
		fmt.Println("Проект с указанным именем уже существует!")
		os.Exit(1)
	} else { //Создание дирректории для проекта
		err := os.MkdirAll(pathToProject, 0755)
		if err != nil {
			panic(err)
		}
	}

	//Генерация файла service
	s := systemd.Service{NameProject: (*args)[2]}
	s.Create()
}

func existDir(path string) bool {
	//Функция проверяет, существует ли данная дирректория
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}

}

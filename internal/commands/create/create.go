package create

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/nginx"
	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/internal/commands/systemd"
	"nginx-gunicorn-systemctl/pkg/osdir"
	"os"
	"os/user"
	"strconv"
)

const pathToProdPjt = "/opt/ngs/production/"

var pathToProject string

func Create(args *[]string) {
	//Проверка задано ли имя проекта после команды create
	if len(*args) < 3 {
		fmt.Println("Не указано имя проекта!")
		os.Exit(0)
	}

	pathToProject = pathToProdPjt + (*args)[2] //Проверка на уже существующий проект

	if osdir.ExistDir(pathToProject) {
		//Проверка на уже существующий проект
		fmt.Println("Проект с указанным именем уже существует!")
		os.Exit(1)
	} else {
		//Создание директории для проекта
		osdir.CreateAllDir(pathToProject)

		//Добавление группы ngs в группу доступа каталога
		gid, err := user.LookupGroup("ngs") //Получение gid группы
		if err == nil {
			gidtoint, err := strconv.Atoi(gid.Gid) //Конвертирование gid в int
			if err != nil {
				panic(err)
			}

			os.Chown(pathToProject, os.Getuid(), gidtoint) // Установка группы доступа каталога
			err = os.Chmod(pathToProject, 0770)            //Установка прав 0700 на каталог
			if err != nil {
				panic(err)
			}
		}

		//Генерация файла service
		service := systemd.Service{NameProject: (*args)[2]}
		service.Create()

		//Генерация файла socket
		socket := systemd.Socket{NameProject: (*args)[2]}
		socket.Create()

		//Запуск команды demon-reload
		systemctl.DaemonReload()

		//Создание конфигурационных файлов nginx
		nginx.Nginx{NameProject: (*args)[2]}.CreateConf()
	}

}

package create

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"nginx-gunicorn-systemctl/internal/commands/nginx"
	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/internal/commands/systemd"
	"nginx-gunicorn-systemctl/internal/conf"
	"nginx-gunicorn-systemctl/pkg/filedirmanager"
	fmd "nginx-gunicorn-systemctl/pkg/filedirmanager"
)

func CreateTest(args *[]string) {
	//Проверка задано ли имя проекта после команды create
	if len(*args) < 3 {
		fmt.Println("Не указано имя проекта!")
		os.Exit(0)
	}

	pathToPjt := conf.PathToDevPjt + (*args)[2]

	//Проверка на уже существующий проект
	if filedirmanager.ExistDir(pathToPjt) {
		//Проверка на уже существующий проект
		fmt.Println("Проект с указанным именем уже существует!")
		os.Exit(1)
	} else {
		//Создание директории для проекта
		filedirmanager.CreateAllDir(pathToPjt)

		//Добавление группы ngs в группу доступа каталога
		gid, err := user.LookupGroup("ngs") //Получение gid группы
		if err == nil {
			gidtoint, err := strconv.Atoi(gid.Gid) //Конвертирование gid в int
			if err != nil {
				panic(err)
			}

			os.Chown(pathToPjt, os.Getuid(), gidtoint) // Установка группы доступа каталога
			err = os.Chmod(pathToPjt, 0770)            //Установка прав 0700 на каталог
			if err != nil {
				panic(err)
			}
		}

		//Генерация файла service
		service := systemd.Service{NameProject: (*args)[2], Production: false}
		service.Create()

		//Генерация файла socket
		socket := systemd.Socket{NameProject: (*args)[2], Production: false}
		socket.Create()

		//Запуск команды demon-reload
		systemctl.DaemonReload()

		//Создание конфигурационных файлов nginx
		nginx.Nginx{NameProject: (*args)[2], Production: false}.CreateConf()

		//Добавление доменной записи для тестов в /etc/hosts
		fmd.AddStringEnd(conf.PathToHosts, "127.0.0.1 "+(*args)[2]+"."+conf.DomainDevelopmen)
	}
}

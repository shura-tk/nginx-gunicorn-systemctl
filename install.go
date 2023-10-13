package main

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"

	"nginx-gunicorn-systemctl/internal/conf"
	"nginx-gunicorn-systemctl/pkg/filedirmanager"
)

func main() {
	// 1. Копирование ~cmd/ngs/ngs в /bin
	cmd := "cp cmd/ngs/ngs /bin"
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	// 2. Создание дирректорий /etc/ngs/templates и копирование
	if !filedirmanager.ExistDir(conf.PathToTemplates) {
		filedirmanager.CreateAllDir(conf.PathToTemplates)
	}

	// 3. Копирование файлов из ~/conf/templates в /etc/ngs/templates
	cmd = "cp -R configs/templates/ " + conf.PathToNgs
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	// 4. Создание каталога для проектов
	if !filedirmanager.ExistDir(conf.PathToDevPjt) {
		filedirmanager.CreateAllDir(conf.PathToDevPjt)
	}

	if !filedirmanager.ExistDir(conf.PathToProdPjt) {
		filedirmanager.CreateAllDir(conf.PathToProdPjt)
	}

	//5. Создать группу пользователей, имеющие доступ к /opt/ngs
	cmd = "sudo groupadd ngs -f"
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	//Добавление группы ngs в группу доступа каталога
	gid, err := user.LookupGroup("ngs") //Получение gid группы
	if err == nil {
		gidtoint, err := strconv.Atoi(gid.Gid) //Конвертирование gid в int
		if err != nil {
			panic(err)
		}

		os.Chown(conf.PathToProjects, os.Getuid(), gidtoint) // Установка группы доступа каталога
		err = os.Chmod(conf.PathToProjects, 0770)            //Установка прав 0700 на каталог
		if err != nil {
			panic(err)
		}

		os.Chown(conf.PathToDevPjt, os.Getuid(), gidtoint) // Установка группы доступа каталога
		err = os.Chmod(conf.PathToDevPjt, 0770)            //Установка прав 0700 на каталог
		if err != nil {
			panic(err)
		}

		os.Chown(conf.PathToProdPjt, os.Getuid(), gidtoint) // Установка группы доступа каталога
		err = os.Chmod(conf.PathToProdPjt, 0770)            //Установка прав 0700 на каталог
		if err != nil {
			panic(err)
		}
	}

}

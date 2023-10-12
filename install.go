package main

import (
	"nginx-gunicorn-systemctl/pkg/osdir"
	"os/exec"
)

const pathToTemplates = "/etc/ngs/templates/"
const pathToNgs = "/etc/ngs/"                //Путь до каталога с проектами
const pathToProdPjt = "/opt/ngs/production/" //Путь для продакшин проектов
const pathToDevPjt = "/opt/ngs/development/" //Путь для разрабатываемых и тестируемых проектов
const pathToService = "/etc/systemd/system/"

func main() {
	// 1. Копирование ~cmd/ngs/ngs в /bin
	cmd := "cp cmd/ngs/ngs /bin"
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	// 2. Создание дирректорий /etc/ngs/templates и копирование
	if !osdir.ExistDir(pathToTemplates) {
		osdir.CreateAllDir(pathToTemplates)
	}

	// 3. Копирование файлов из ~/configs/templates в /etc/ngs/templates
	cmd = "cp -R configs/templates/ " + pathToNgs
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	// 4. Создание каталога для проектов
	if !osdir.ExistDir(pathToDevPjt) {
		osdir.CreateAllDir(pathToDevPjt)
	}

	if !osdir.ExistDir(pathToProdPjt) {
		osdir.CreateAllDir(pathToProdPjt)
	}

	//5. Создать группу пользователей, имеющие доступ к /opt/ngs
	cmd = "sudo groupadd ngs -f"
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

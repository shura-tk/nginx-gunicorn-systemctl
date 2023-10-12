package systemd

import (
	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"os"
	"text/template"
)

const pathToService = "/etc/systemd/system/"
const pathToSocket = pathToService
const pathToTemplates = "/etc/ngs/templates/"

type Service struct {
	NameProject string
}

type Socket struct {
	NameProject string
}

func (s *Socket) Create() {
	s.generateFile()
	systemctl.DaemonReload()
}

func (s *Socket) generateFile() {
	//Генерация socket-файла на основе шаблона и сохранение его.

	tmp, err := template.New("socket").ParseFiles(pathToTemplates + "socket")
	if err != nil {
		panic(err)
	}

	//Создание файла
	file, err := os.Create(pathToService + "ngs_" + s.NameProject + ".socket")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	//Вставка значений в шаблон
	err = tmp.Execute(file, s)

	if err != nil {
		panic(err)
	}

}

func (s *Service) Create() {
	s.generateFile()
	systemctl.DaemonReload()

}

func (s *Service) generateFile() {
	//Генерация systend-файла на основе шаблона и сохранение его.

	tmp, err := template.ParseFiles(pathToTemplates + "service")
	if err != nil {
		panic(err)
	}

	//Создание файла
	file, err := os.Create(pathToService + "ngs_" + s.NameProject + ".service")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	//Вставка значений в шаблон
	err = tmp.Execute(file, s)

	if err != nil {
		panic(err)
	}

}

package systemd

import (
	"os"
	"text/template"
)

const pathToService = "/etc/systemd/system/"
const pathToSocket = pathToService
const pathToTemplates = "/etc/ngs/templates/service"

type Service struct {
	NameProject string
}

type socket struct {
	nameProject string
}

func (s *Service) Create() {
	s.generateFile()

}

func (s Service) generateFile() {
	//Генерация systend-файла на основе шаблона и сохранение его.

	tmp, err := template.New("service").ParseFiles(pathToTemplates)
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

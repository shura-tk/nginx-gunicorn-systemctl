package systemd

import (
	"os"
	"text/template"

	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/internal/conf"
)

type Service struct {
	NameProject string
	Production  bool
	Prefix      string
	PathPjt     string
}

type Socket struct {
	NameProject string
	Production  bool
	Prefix      string
	PathPjt     string
}

func (s *Socket) Create() {
	if !s.Production {
		s.Prefix = conf.PrefixDevelopment
		s.PathPjt = conf.PathToDevPjt
	} else {
		s.Prefix = conf.PrefixProduction
		s.PathPjt = conf.PathToProdPjt
	}

	s.generateFile()
	systemctl.DaemonReload()
}

func (s *Socket) generateFile() {
	//Генерация socket-файла на основе шаблона и сохранение его.

	tmp, err := template.New("socket").ParseFiles(conf.PathToTemplates + "socket")
	if err != nil {
		panic(err)
	}

	//Создание файла
	file, err := os.Create(conf.PathToService + s.Prefix + s.NameProject + ".socket")
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
	if !s.Production {
		s.Prefix = conf.PrefixDevelopment
		s.PathPjt = conf.PathToDevPjt
	} else {
		s.Prefix = conf.PrefixProduction
		s.PathPjt = conf.PathToProdPjt
	}
	s.generateFile()
	systemctl.DaemonReload()

}

func (s *Service) generateFile() {
	//Генерация systend-файла на основе шаблона и сохранение его.

	tmp, err := template.ParseFiles(conf.PathToTemplates + "service")
	if err != nil {
		panic(err)
	}

	//Создание файла
	file, err := os.Create(conf.PathToService + s.Prefix + s.NameProject + ".service")
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

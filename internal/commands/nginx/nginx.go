package nginx

import (
	"fmt"
	"os"
	"text/template"

	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/internal/conf"
	"nginx-gunicorn-systemctl/pkg/filedirmanager"
)

type Nginx struct {
	NameProject string
	Domain      string
	Production  bool
	Prefix      string
	PathPjt     string
}

//const defPathNginxConf = "/etc/nginx/sites-available/"

func (n Nginx) CreateConf() {
	//Задание значения префикса
	if !n.Production {
		n.Prefix = conf.PrefixDevelopment
		n.PathPjt = conf.PathToDevPjt
		n.Domain = conf.DomainDevelopmen
	} else {
		n.Prefix = conf.PrefixProduction
		n.PathPjt = conf.PathToProdPjt
		n.Domain = conf.DomainProduction
	}

	n.generationConf()
	n.createSimlink()

	// Перезагрузка службы nginx
	n.serviceReload()
}

func (n Nginx) createSimlink() {
	filedirmanager.CreateAllDir(conf.PathToNginxConfEnabled)
	os.Symlink(conf.PathToNginxConfAvailable+n.Prefix+n.NameProject, conf.PathToNginxConfEnabled+n.Prefix+n.NameProject)
}

func (n Nginx) serviceReload() {
	sctl := systemctl.New(n.Prefix + n.NameProject)
	sctl.Restart()
}

func (n Nginx) generationConf() {
	_, err := os.Stat(conf.PathToNginxConfAvailable + n.Prefix + n.NameProject) //Проверка на наличие файла по указанному пути
	if err != nil {                                                             //Если не удается найти указанный файл по пути
		if os.IsNotExist(err) {

			filedirmanager.CreateAllDir(conf.PathToNginxConfAvailable)

			//Создание файла конфигурации
			file, err := os.Create(conf.PathToNginxConfAvailable + n.Prefix + n.NameProject)
			if err != nil {
				panic(err)
			}

			temp, err := template.ParseFiles(conf.PathToTemplatesNginx)
			if err != nil {
				fmt.Println(err)
			}

			//Вставка в файл конфигурации настроек из шаблона с подстановокой данных
			err = temp.Execute(file, n)
			if err != nil {
				fmt.Println(err)
			}

			err = file.Close()
			if err != nil {
				panic(err)
			}
		}
	}
}

package nginx

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/pkg/osdir"
	"os"
	"text/template"
)

type Nginx struct {
	NameProject string
}

const defPathNginxConf = "/etc/nginx/sites-available/"

func (n Nginx) CreateConf() {
	n.generationConf()
	n.createSimlink()

	// Перезагрузка службы nginx
	n.serviceReload()

}

func (n Nginx) createSimlink() {
	path := "/etc/nginx/sites-enabled/"
	osdir.CreateAllDir(path)
	os.Symlink(defPathNginxConf+n.NameProject, path+n.NameProject)
}

func (n Nginx) serviceReload() {
	sctl := systemctl.New("ngs_" + n.NameProject)
	sctl.Restart()
}

func (n Nginx) generationConf() {

	_, err := os.Stat(defPathNginxConf + n.NameProject) //Проверка на наличие файла по указанному пути
	if err != nil {                                     //Если не удается найти указанный файл по пути
		if os.IsNotExist(err) {

			osdir.CreateAllDir(defPathNginxConf)

			//Создание файла конфигурации
			file, err := os.Create(defPathNginxConf + n.NameProject)
			if err != nil {
				panic(err)
			}

			temp, err := template.ParseFiles("/etc/ngs/templates/nginx")
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

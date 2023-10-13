package restart

import (
	"fmt"
	"os"

	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/internal/conf"
	"nginx-gunicorn-systemctl/pkg/filedirmanager"
)

func RestartTest(args *[]string) {
	//Проверка задано ли имя проекта после команды restart
	if len(*args) < 3 {
		fmt.Println("Не указано имя проекта, который необходимо перезапустить!")
		os.Exit(0)
	}

	if !filedirmanager.ExistDir(conf.PathToDevPjt + (*args)[2]) { //Проверка на отсутствие проекта
		fmt.Println("Проект или служба с указанным именем не найден!")
		os.Exit(1)
	} else {
		sctl := systemctl.New(conf.PrefixDevelopment + (*args)[2])
		sctl.Restart()
	}

}

package restart

import (
	"fmt"
	"nginx-gunicorn-systemctl/internal/commands/systemctl"
	"nginx-gunicorn-systemctl/pkg/osdir"
	"os"
)

func Restart(args *[]string) {
	//Проверка задано ли имя проекта после команды restart
	if len(*args) < 3 {
		fmt.Println("Не указано имя проекта, который необходимо перезапустить!")
		os.Exit(0)
	}

	if !osdir.ExistDir("/opt/ngs/development/" + (*args)[2]) { //Проверка на отсутствие проекта
		fmt.Println("Проект или служба с указанным именем не найден!")
		os.Exit(1)
	} else {
		sctl := systemctl.New((*args)[2])
		sctl.Restart()
	}

}

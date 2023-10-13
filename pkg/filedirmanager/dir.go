package filedirmanager

import (
	"os"
)

func CreateAllDir(path string) {

	err := os.MkdirAll(path, 0770)
	if err != nil {
		panic(err)
	}
}

func ExistDir(path string) bool {
	//Функция проверяет, существует ли данная дирректория
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}

}

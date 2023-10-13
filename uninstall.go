package main

import "os/exec"

func main() {
	cmd := "sudo rm /opt/ngs/ -rf && sudo rm /etc/ngs/ -rf && sudo rm /bin/ngs && sudo rm /etc/systemd/system/ngs/ -rf"
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	//Удаление из /etc/systemd/system
	cmd = "sudo rm /etc/systemd/system/ngs_*.socket && sudo rm /etc/systemd/system/ngs_*.service"
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	//Удаление из nginx
	cmd = "sudo rm /etc/nginx/sites-available/ngs_* && sudo rm /etc/nginx/sites-enabled/ngs_*"
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

	//Удаление группы пользователей
	cmd = "sudo groupdell ngs"
	err = exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}

}

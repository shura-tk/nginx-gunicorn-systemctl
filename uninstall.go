package main

import "os/exec"

func main() {
	cmd := "sudo rm /opt/ngs/ -rf && sudo rm /etc/ngs/ -rf && sudo rm /bin/ngs && sudo rm /etc/systemd/system/ngs/ -rf"
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

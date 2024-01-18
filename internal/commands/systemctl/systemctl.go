package systemctl

import "os/exec"

type systemctl struct {
	name string
}

func New(name string) *systemctl {
	return &systemctl{name: name}
}

func (s *systemctl) Restart() {
	cmd := "systemctl restart ngs_" + s.name
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

func (s *systemctl) Stop() {
	cmd := "systemctl stop ngs_" + s.name
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

func (s *systemctl) Start() {
	cmd := "systemctl start ngs_" + s.name
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

func DaemonReload() {
	cmd := "systemctl daemon-reload"
	err := exec.Command("bash", "-c", cmd).Start()
	if err != nil {
		panic(err)
	}
}

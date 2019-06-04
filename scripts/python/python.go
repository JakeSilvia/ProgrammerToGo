package scripts

import (
	"github.com/ProgrammerToGo/scripts"
	"log"
	"strings"
	"fmt"
)

type Python struct {
	PipVersion float32
	Libraries []string
}

type PythonUtility interface {
	GetLibraries() ([]string, error)
	ScanPip() error
	InstallPip()
	InstallLibraries() error
}

const (
	pythonLibCmd  = `pip freeze`
	pythonInstallLibCmd  = `pip install`
	pipInstallCmd = `sudo easy_install pip`
	pipInstallAlt = `curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && sudo python get-pip.py`
	pipScanCmd = `pip -V`
	couldNotUninstall = `sudo pip install --ignore-installed `
)

func GetIgnoreInstallCommand(lib string) (string) {
	return couldNotUninstall + lib
}

func (py *Python) ParseLibraries() (error) {
	command, err := scripts.RunCommand(pythonLibCmd)
	if err != nil {
		return err
	}

	py.Libraries = strings.Split(command.Output, "\n")
	return nil
}

func (py *Python) ScanPip() (error) {
	out, err := scripts.RunCommand(pipScanCmd)
	if err != nil {
		log.Printf("error installing PIP: %v", err)
		return err
	}
	fmt.Printf("pip: %v", out.Output)

	return nil
}

func (py *Python) InstallPip() error {
	out, err := scripts.RunCommand(pipInstallCmd)
	if err != nil {
		log.Printf("error installing PIP: %v", out.Output)
		return err
	}

	return nil
}

func (py *Python) InstallLibraries() error {
	installCmd := pythonInstallLibCmd
	for _, lib := range py.Libraries{
		installCmd += lib + " "
	}

	out, err := scripts.RunCommand(installCmd)
	if err != nil {
		log.Printf("error installing PIP: %v", out.Output)
		return err
	}

	return nil
}

func NewPython () *Python {
	return &Python{}
}
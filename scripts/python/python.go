package scripts

import (
	"github.com/goSwap/scripts"
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
	ScanPip() (error)
	InstallPip()
}

const (
	pythonLibCmd  = `pip freeze`
	pipInstallCmd = `sudo easy_install pip`
	pipScanCmd = `pip -V`
)

func (py *Python) ParseLibraries() (error) {
	out, err := scripts.GetCommandOutput(pythonLibCmd)
	if err != nil {
		return err
	}

	py.Libraries = strings.Split(out, "\n")
	return nil
}

func (py *Python) ScanPip() (error) {
	out, err := scripts.GetCommandOutput(pipScanCmd)
	if err != nil {
		log.Printf("error installing PIP: %v", err)
		return err
	}
	fmt.Printf("pip: %v", out)

	return nil
}

func (py *Python) InstallPip() error {
	out, err := scripts.GetCommandOutput(pipInstallCmd)
	if err != nil {
		log.Printf("error installing PIP: %v", out)
		return err
	}

	return nil
}

func NewPython () *Python {
	return &Python{}
}
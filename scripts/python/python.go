package scripts

import (
	"github.com/goSwap/scripts"
	"strings"
)

const pythonLibCmd = `pip freeze`

func GetLibraries() ([]string, error) {
	out, err := scripts.RunCommand(pythonLibCmd)
	if err != nil{
		return []string{}, nil
	}

	return strings.Split(out, "\n"), nil
}
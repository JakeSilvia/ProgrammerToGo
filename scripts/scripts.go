package scripts

import (
	"strings"
	"os/exec"
	"log"
)

type Command struct {
	exitCode int
	output string
}

func GetCommandOutput(command string) (string, error) {
	cmds := strings.Split(command, "\n")
	var output string
	for _, iCmd := range cmds {
		out, err := exec.Command("/bin/sh", "-c", iCmd).Output()
		if err != nil {
			log.Printf("error running command: %s\n%s", iCmd, err)
			return output, err
		}
		output += string(out)
	}

	return output, nil
}

func GetCommandStatus(command string) (error) {
	err := exec.Command("/bin/sh", "-c", command).Run()
	if err != nil {
		log.Printf("error running command: %s\n%s", command, err)
		return err
	}

	return nil
}

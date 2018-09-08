package scripts

import (
	"strings"
	"os/exec"
	"log"
)

func RunCommand(command string) (string, error) {
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

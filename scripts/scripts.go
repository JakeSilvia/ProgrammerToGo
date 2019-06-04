package scripts

import (
	"os/exec"
	"log"
	"bufio"
	"syscall"
)

type Command struct {
	ExitCode int
	Output   string
}

const (
	brewInstall = `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
)

func RunCommand(command string) (*Command, error) {
	Cmd := &Command{}
	proc := exec.Command("/bin/sh", "-c", command)
	cmdReader, err := proc.StdoutPipe()
	if err != nil {
		log.Printf("error creating stdout for command: %s\n", err)
	}

	cmdError, err := proc.StderrPipe()
	if err != nil {
		log.Printf("error creating stderr for command: %s\n", err)
	}

	scannerStdout := bufio.NewScanner(cmdReader)
	scannerStderr := bufio.NewScanner(cmdError)
	go func() {
		go func() {
			for scannerStdout.Scan() {
				out := scannerStdout.Text()
				log.Printf("Command (stdout): %s\n", out)
				Cmd.Output += out
			}
		}()

		for scannerStderr.Scan() {
			out := scannerStderr.Text()
			log.Printf("Command (stderr): %s\n", out)
			Cmd.Output += out
		}
	}()

	err = proc.Run()
	if err != nil {
		log.Printf("Command Error: %s\n", err)
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			Cmd.ExitCode = ws.ExitStatus()
		}
	} else {
		// success, ExitCode should be 0 if go is ok
		ws := proc.ProcessState.Sys().(syscall.WaitStatus)
		Cmd.ExitCode = ws.ExitStatus()
	}
	return Cmd, nil
}

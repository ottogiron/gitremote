package server

import (
	"bufio"
	"strings"

	"os/exec"

	"github.com/pkg/errors"
)

const (
	successStatus = 0
	failStatus    = 1
)

//GitService operations in the git server
type GitService interface {
	Execute(command string, onOutput func(string)) error
}

var _ GitService = (*gitService)(nil)

type gitService struct {
}

func (g *gitService) Execute(command string, onOutput func(string)) error {

	commandTokens := strings.Fields(command)

	if len(command) == 0 {
		return errors.New("Invalid empty command")
	}

	if commandTokens[0] != "git" {
		return errors.Errorf("Invalid non git command: %s", command)
	}

	cmdName := "git"
	cmdArgs := commandTokens[1:]

	cmd := exec.Command(cmdName, cmdArgs...)

	cmdReader, err := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			onOutput(scanner.Text())
		}
	}()

	if err != nil {
		return errors.Wrapf(err, "Errorr creating StdoutPipe for command %s", command)
	}

	err = cmd.Start()

	if err != nil {
		return errors.Wrapf(err, "Errors starting command %s", command)
	}

	err = cmd.Wait()

	if err != nil {
		return errors.Wrapf(err, "Error waiting for command %s", command)
	}

	return nil
}

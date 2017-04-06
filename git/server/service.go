package server

import (
	"bufio"
	"strings"

	"os/exec"

	"sync"

	"github.com/pkg/errors"
)

const (
	successStatus = 0
	failStatus    = 1
)

//GitService operations in the git server
type GitService interface {
	Execute(dir string, command string, onOutput func(string)) error
}

var _ GitService = (*gitService)(nil)

type gitService struct {
	mu sync.Mutex
}

//NewGitService returns a new GitService
func NewGitService() GitService {
	return &gitService{}
}

func (g *gitService) Execute(dir string, command string, onOutput func(string)) error {

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
	cmd.Dir = dir
	cmdReader, err := cmd.StdoutPipe()

	scanner := bufio.NewScanner(cmdReader)
	go func() {

		for {
			g.mu.Lock()
			token := scanner.Scan()
			g.mu.Unlock()
			if !token {
				return
			}
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
	g.mu.Lock()
	err = cmd.Wait()
	g.mu.Unlock()
	if err != nil {
		return errors.Wrapf(err, "Error waiting for command %s", command)
	}

	return nil
}

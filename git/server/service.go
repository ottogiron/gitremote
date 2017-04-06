package server

import (
	"strings"

	"os/exec"

	"sync"

	"io"

	"github.com/pkg/errors"
)

const (
	successStatus = 0
	failStatus    = 1
)

//GitService operations in the git server
type GitService interface {
	Execute(dir string, command string, output io.Writer) error
}

var _ GitService = (*gitService)(nil)

type gitService struct {
	mu sync.Mutex
}

//NewGitService returns a new GitService
func NewGitService() GitService {
	return &gitService{}
}

func (g *gitService) Execute(dir string, command string, output io.Writer) error {

	commandTokens := strings.Fields(command)

	if len(command) == 0 {
		return errors.Errorf("Invalid empty command %s", command)
	}

	if commandTokens[0] != "git" {
		return errors.Errorf("Invalid non git command: %s", command)
	}

	cmdName := "git"
	cmdArgs := commandTokens[1:]

	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Dir = dir
	cmd.Stdout = output

	err := cmd.Run()

	if err != nil {
		return errors.Wrapf(err, "Errors runing command %s", command)
	}
	return nil
}

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
	mu                 sync.Mutex
	allowedDirectories []string
	allowedCommands    []string
}

//NewGitService returns a new GitService
func NewGitService(options ...Option) GitService {
	g := &gitService{}
	for _, option := range options {
		option(g)
	}
	return g
}

func (g *gitService) Execute(dir string, command string, output io.Writer) error {
	if !g.isAllowedDirectory(dir) {
		return errors.Errorf("Directory  is not allowed %s", dir)
	}

	if !g.isAllowedCommand(command) {
		return errors.Errorf("Executed command is not allowed %s", command)
	}

	command = strings.Replace(command, "\\", "", -1)

	if len(command) == 0 {
		return errors.Errorf("Invalid empty command %s", command)
	}

	cmd := exec.Command("sh", "-c", command)

	cmd.Dir = dir
	cmd.Stdout = output
	cmd.Stderr = output
	err := cmd.Run()

	if err != nil {
		return errors.Wrapf(err, "Failed to run command %s", command)
	}
	return nil
}

func (g *gitService) isAllowedDirectory(directory string) bool {
	for _, d := range g.allowedDirectories {
		if directory == d {
			return true
		}
	}
	return false
}

func (g *gitService) isAllowedCommand(command string) bool {

	if len(command) == 0 {
		return false
	}

	commandFields := strings.Fields(command)

	if commandFields[0] != "git" {
		return false
	}

	for _, cmd := range g.allowedCommands {
		allowedCmdFields := strings.Fields(cmd)
		if len(commandFields) == len(allowedCmdFields) {
			valid := true
			for i, cmdAllowedField := range allowedCmdFields {
				valid = valid && (cmdAllowedField == commandFields[i])
			}
			if valid {
				return true
			}
		}
	}

	return false
}

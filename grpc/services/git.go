package services

import (
	"io"
	"os"

	"github.com/inconshreveable/log15"
	"github.com/ottogiron/gitremote/git/server"
	"github.com/ottogiron/gitremote/grpc/gen"
	"github.com/pkg/errors"
)

var _ gen.GitServiceServer = (*GitServiceServer)(nil)

//GitServiceServer implementation of rpc git  client service
type GitServiceServer struct {
	service server.GitService
}

//NewGitServiceServer returns a new git service
func NewGitServiceServer(service server.GitService) *GitServiceServer {
	return &GitServiceServer{service}
}

type streamWriter struct {
	stream gen.GitService_ExecuteServer
}

func (s *streamWriter) Write(p []byte) (int, error) {
	err := s.stream.Send(&gen.Output{
		Message: string(p),
	})
	if err != nil {
		return 0, errors.Wrapf(err, "Failed to write output to stream %s", err)
	}
	return len(p), nil
}

//Execute executes a git command in the remote server
func (g *GitServiceServer) Execute(command *gen.Command, stream gen.GitService_ExecuteServer) error {

	log15.Info("Executing git command", "cmd", command.Command, "dir", command.Dir)
	stdout := io.MultiWriter(os.Stdout, &streamWriter{stream})
	err := g.service.Execute(command.Dir, command.Command, stdout)

	if err != nil {
		log15.Error("Failed to execute git command", "err", err)
		return errors.Wrapf(err, "Failed to execute git command in the server", command.Command)
	}

	return nil
}

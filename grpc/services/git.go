package services

import (
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

//Execute executes a git command in the remote server
func (g *GitServiceServer) Execute(command *gen.Command, stream gen.GitService_ExecuteServer) error {

	err := g.service.Execute(command.Command, func(msgOutput string) {
		log15.Info(msgOutput)
		stream.Send(&gen.Output{
			Message: msgOutput,
		})
	})

	if err != nil {
		errors.Wrapf(err, "Failed to execute git command in the server", command.Command)
	}
	return nil
}

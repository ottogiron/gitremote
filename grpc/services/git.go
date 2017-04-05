package services

import "github.com/ottogiron/gitremote/grpc/gen"

var _ gen.GitServiceServer = (*GitService)(nil)

//GitService implementation of rpc git  client service
type GitService struct {
}

//NewGitService returns a new git service
func NewGitService() *GitService {
	return &GitService{}
}

//Execute executes a git command in the remote server
func (g *GitService) Execute(command *gen.Command, stream gen.GitService_ExecuteServer) error {

	return nil
}

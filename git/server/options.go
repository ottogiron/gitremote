package server

//Option option for the git service
type Option func(*gitService)

//SetAllowedDirectories set service allowed directories
func SetAllowedDirectories(directories []string) Option {
	return func(g *gitService) {
		g.allowedDirectories = directories
	}
}

//SetAllowedCommands set service allowed commands
func SetAllowedCommands(commands []string) Option {
	return func(g *gitService) {
		g.allowedCommands = commands
	}
}

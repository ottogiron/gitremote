package cmd

import (
	"context"
	"fmt"

	"io"

	"github.com/inconshreveable/log15"
	"github.com/ottogiron/gitremote/grpc/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	hostAddressKey = "host-address"
	workingDirKey  = "working-dir"
	commandKey     = "command"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes a git command in a remote server",
	Long: `Executes a git command in a  remote server:

Example:.
gitr exec --command="git status" --working-dir=/home/otto/myproject --host-address=myhost:2183

The above executes translates to "git status" command running in the /home/otto/myproject directory
`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString(hostAddressKey)
		dir := viper.GetString(workingDirKey)
		command := viper.GetString(commandKey)

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log15.Crit("did not connect to remote server ", "err", err)
			return
		}
		defer conn.Close()
		gitClient := gen.NewGitServiceClient(conn)

		rstream, err := gitClient.Execute(context.Background(), &gen.Command{Dir: dir, Command: command})

		if err != nil {
			log15.Crit("could not execute git command", "err", err)
			return
		}

		for {
			out, err := rstream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				//gprclog.Fatalf()
			}

			fmt.Print(out.Message)
		}
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	execCmd.Flags().StringP(hostAddressKey, "a", "localhost:2183", "git remote server address")
	execCmd.Flags().StringP(workingDirKey, "w", ".", "Working directory of the git command to run")
	execCmd.Flags().StringP(commandKey, "c", "", "git command to be executed example: git status")

	viper.BindPFlag(hostAddressKey, execCmd.Flags().Lookup(hostAddressKey))
	viper.BindPFlag(workingDirKey, execCmd.Flags().Lookup(workingDirKey))
	viper.BindPFlag(commandKey, execCmd.Flags().Lookup(commandKey))

}

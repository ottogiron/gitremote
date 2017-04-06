package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		// TODO: Work your own magic here
		fmt.Println("exec called")
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

}

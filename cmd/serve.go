package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/inconshreveable/log15"
	"github.com/ottogiron/gitremote/git/server"
	"github.com/ottogiron/gitremote/grpc/services"

	"github.com/ottogiron/gitremote/grpc/gen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	portKey               = "port"
	tlsKey                = "tls"
	certFileKey           = "cert-file"
	keyFileKey            = "key-file"
	allowedDirectoriesKey = "allowed-directories"
	allowedCommandsKey    = "allowed-commands"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the git remote server",
	Long: `Starts the git remote server

Example:.
gitr serve --port=2183`,
	Run: func(cmd *cobra.Command, args []string) {
		tls := viper.GetBool(tlsKey)
		certFile := viper.GetString(certFileKey)
		keyFile := viper.GetString(keyFileKey)
		allowedDirectories := viper.GetStringSlice(allowedDirectoriesKey)
		allowedCommands := viper.GetStringSlice(allowedCommandsKey)

		var opts []grpc.ServerOption
		if tls {
			creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
			if err != nil {
				grpclog.Fatalf("Failed to generate credentials %v", err)
				return
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}
		grpcServer := grpc.NewServer(opts...)

		_, cancel := context.WithCancel(context.Background())
		defer cancel()

		gitService := server.NewGitService(
			server.SetAllowedCommands(allowedCommands),
			server.SetAllowedDirectories(allowedDirectories),
		)

		gitServerService := services.NewGitServiceServer(gitService)

		gen.RegisterGitServiceServer(grpcServer, gitServerService)

		port := viper.GetInt(portKey)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

		if err != nil {
			log15.Crit("FAiled to create a grpc server", "err", err)
			return
		}

		log15.Info("Listening...	", "port", port)
		log15.Info("Allowed directories", "list", allowedDirectories)
		log15.Info("Allowed commands", "list", allowedCommands)
		err = grpcServer.Serve(lis)
		if err != nil {
			log15.Crit("Failed to start grpc server", "err", err)
			return

		}

	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP(portKey, "p", 2183, "Port for the rpc service")
	serveCmd.Flags().BoolP(tlsKey, "t", false, "Connection uses TLS if true, else plain TCP")
	serveCmd.Flags().StringP(certFileKey, "c", "server.pem", "The TLS cert file")
	serveCmd.Flags().StringP(keyFileKey, "k", "server.key", "The TLS key file")
	serveCmd.Flags().StringSliceP(allowedDirectoriesKey, "d", []string{}, "The list of allowed git executable directories")
	serveCmd.Flags().StringSliceP(allowedCommandsKey, "a", []string{}, "The list of allowed git commands")

	viper.BindPFlag(portKey, serveCmd.Flags().Lookup(portKey))
	viper.BindPFlag(tlsKey, serveCmd.Flags().Lookup(tlsKey))
	viper.BindPFlag(certFileKey, serveCmd.Flags().Lookup(certFileKey))
	viper.BindPFlag(keyFileKey, serveCmd.Flags().Lookup(keyFileKey))
	viper.BindPFlag(allowedDirectoriesKey, serveCmd.Flags().Lookup(allowedDirectoriesKey))
	viper.BindPFlag(allowedCommandsKey, serveCmd.Flags().Lookup(allowedCommandsKey))

}

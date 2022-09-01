package main

import (
	"os"
	"strconv"

	"github.com/WHGLD/antibruteforce/internal/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var rootCmd = &cobra.Command{
	Use:  "cli <command> <params>",
	Long: `CLI interface for ABF service`,
}

var (
	host string
	port int
)

func main() {
	rootCmd.PersistentFlags().StringVar(
		&host,
		"host",
		"localhost",
		"gRPC server host (default 'localhost')",
	)
	rootCmd.PersistentFlags().IntVarP(
		&port,
		"port",
		"p",
		50051,
		"gRPC server port (default '50051')",
	)

	if rootCmd.Execute() != nil {
		os.Exit(1) //nolint:gocritic
	}
}

func getGRPCClient() (server.ABruteforceClient, error) {
	conn, err := grpc.Dial(
		host+":"+strconv.Itoa(port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return server.NewABruteforceClient(conn), nil
}

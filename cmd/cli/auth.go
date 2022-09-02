package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/WHGLD/antibruteforce/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:  "auth <login> <pass> <ip>",
	Long: "Auth command - check for given login pw and ip",
	Args: cobra.ExactArgs(3),
	Run:  Auth,
}

func Auth(cmd *cobra.Command, args []string) {
	login := args[0]
	pw := args[1]
	ip := args[2]

	client, err := getGRPCClient()
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(504)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	response, errAuth := client.Auth(ctx, &server.AuthRequest{Login: login, Password: pw, Ip: ip})

	cancel()

	if errAuth != nil {
		fmt.Printf(
			"error while auth [login: %s password: %s ip: %s]: %s\n",
			login,
			pw,
			ip,
			errAuth.Error(),
		)
		os.Exit(13)
	}
	fmt.Printf(
		"auth [login: %s password: %s ip: %s] ok - %t\n",
		login,
		pw,
		ip,
		response.GetOk(),
	)
}

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
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:  "reset <login> <pass> <ip>",
	Long: "Reset command - delete stats for given login and ip",
	Args: cobra.ExactArgs(2),
	Run:  Reset,
}

func Reset(cmd *cobra.Command, args []string) {
	login := args[0]
	pw := args[1]
	ip := args[2]

	// todo проверка ip и mask

	client, err := getGRPCClient()
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(504)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = client.Reset(ctx, &server.ResetRequest{Login: login, Password: pw, Ip: ip})

	cancel()

	if err != nil {
		fmt.Printf(
			"error while resetting stat for [login: %s password: %s ip: %s]: %s\n",
			login,
			pw,
			ip,
			err.Error(),
		)
		os.Exit(13)
	}
	fmt.Printf("stats for [login: %s password: %s ip: %s] successfully reseted\n", login, pw, ip)
}

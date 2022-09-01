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
	rootCmd.AddCommand(removeBlackCmd)
}

var removeBlackCmd = &cobra.Command{
	Use:  "removeBlack <ip> <mask>",
	Long: `Remove from black list command`,
	Args: cobra.ExactArgs(2),
	Run:  removeFromBlackList,
}

func removeFromBlackList(cmd *cobra.Command, args []string) {
	ip := args[0]
	mask := args[1]

	// todo проверка ip и mask

	client, err := getGRPCClient()
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(504)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	request := &server.RemoveNetMaskRequest{Ip: ip, Mask: mask}
	_, err = client.RemoveFromBlackList(ctx, request)

	cancel()

	if err != nil {
		fmt.Printf("error while removing net mask [ip: %s mask: %s]: %s\n", ip, mask, err.Error())
		os.Exit(13)
	}
	fmt.Printf("net mask [ip: %s mask: %s] successfully removed\n", ip, mask)
}

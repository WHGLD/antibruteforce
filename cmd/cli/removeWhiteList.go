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
	rootCmd.AddCommand(removeWhiteCmd)
}

var removeWhiteCmd = &cobra.Command{
	Use:  "removeWhite <ip> <mask>",
	Long: `Remove from white list command`,
	Args: cobra.ExactArgs(2),
	Run:  removeFromWhiteList,
}

func removeFromWhiteList(cmd *cobra.Command, args []string) {
	ip := args[0]
	mask := args[1]

	client, err := getGRPCClient()
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(504)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	request := &server.RemoveNetMaskRequest{Ip: ip, Mask: mask}
	_, err = client.RemoveFromWhiteList(ctx, request)

	cancel()

	if err != nil {
		fmt.Printf("error while removing net mask [ip: %s mask: %s]: %s\n", ip, mask, err.Error())
		os.Exit(13)
	}
	fmt.Printf("net mask [ip: %s mask: %s] successfully removed\n", ip, mask)
}

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
	rootCmd.AddCommand(addWhiteCmd)
}

var addWhiteCmd = &cobra.Command{
	Use:  "addWhite <ip> <mask>",
	Long: `Add to white list command`,
	Args: cobra.ExactArgs(2),
	Run:  AddToWhiteList,
}

func AddToWhiteList(cmd *cobra.Command, args []string) {
	ip := args[0]
	mask := args[1]

	client, err := getGRPCClient()
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(504)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	request := &server.AddNetMaskRequest{Ip: ip, Mask: mask}
	response, errAdd := client.AddToWhiteList(ctx, request)

	cancel()

	if errAdd != nil {
		fmt.Printf("error while adding net mask [ip: %s mask: %s]: %s\n", ip, mask, errAdd.Error())
		os.Exit(13)
	}
	fmt.Printf("net mask [ip: %s mask: %s] successfully added\n", response.GetIp(), response.GetMask())
}

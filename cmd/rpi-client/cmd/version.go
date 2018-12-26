package cmd

import (
	"fmt"
	"os"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client version: %s\n", rpi.Version)
		client, ctx := getGrpcClientAndContext(cmd)
		common := proto.NewCommonClient(client)
		res, err := common.Version(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not retrieve server version: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Server version: %s\n", res.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

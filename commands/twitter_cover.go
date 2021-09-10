package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TwitterCoverCmd = &cobra.Command{
  Use: "cover",
  Short: "publish cover image to profile",
  Run: publishCover,
}

func publishCover(cmd *cobra.Command, args []string) {
  fmt.Println("publish command executed")
}


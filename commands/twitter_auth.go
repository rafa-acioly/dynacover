package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TwitterAuthCmd = &cobra.Command{
  Use: "auth",
  Short: "Validate the twitter auth credentials",
  Long: "Check if the twitter auth credentials are valid making a request to twitter API",
  Run: validateAuth,
}

func validateAuth(cmd *cobra.Command, args []string) {
  fmt.Println("command twitter auth executed.")
}


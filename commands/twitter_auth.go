package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AuthCmd represents the command that will
// verify if the given twitter credentials are working.
var AuthCmd = &cobra.Command{
  Use: "auth",
  Short: "Validate the twitter auth credentials",
  Long: "Check if the twitter auth credentials are valid making a request to twitter API",
  Run: validateAuth,
}

func validateAuth(cmd *cobra.Command, args []string) {
  fmt.Println("command twitter auth executed.")
}


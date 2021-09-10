package main

import (
  "github.com/spf13/cobra"
  "github.com/rafa-acioly/dynacover/commands"
)

var (
  templateFlag string
  rootCmd = &cobra.Command{
    Use: "generate",
    Short: "create twitter cover",
    Long: "create twitter cover and publish into your twitter profile",
  }
)

func main() {
  rootCmd.PersistentFlags().StringVar(&templateFlag, "template", "basic", "choose the cover layout")

  rootCmd.AddCommand(commands.TwitterAuthCmd)
  rootCmd.AddCommand(commands.TwitterCoverCmd)
}


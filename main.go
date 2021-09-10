package main

import (
	"github.com/rafa-acioly/dynacover/commands"
	"github.com/spf13/cobra"
  "log"
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

  rootCmd.AddCommand(commands.CoverCmd)

  log.Fatal(rootCmd.Execute())
}


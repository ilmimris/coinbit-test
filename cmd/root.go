package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "account-svc",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("root command")
	},
}

func Run() {
	rootCommand.Execute()

}

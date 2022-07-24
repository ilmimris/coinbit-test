package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "serve",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Println("rest command")
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("rest command")
	},
	PostRun: func(cmd *cobra.Command, args []string) {},
}

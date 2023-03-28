package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start the plippy service",
		Run: func(cmd *cobra.Command, args []string) {
			err := handleStartCommand(args)

			if err != nil {
				log.Println(err)
			}

		},
	}
}

func handleStartCommand(args []string) error {
	return StartPlippy()
}

func init() {
	startCmd := StartCmd()
	rootCmd.AddCommand(startCmd)
}

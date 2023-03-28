package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list entries",
		Run: func(cmd *cobra.Command, args []string) {
			err := handleListCommand(args)

			if err != nil {
				fmt.Println(err)
			}

		},
	}
}

func handleListCommand(args []string) error {

	entries, err := List()
	if err != nil {
		return err
	}

	selection, err := ShowSelectionView(*entries)
	if err != nil {
		return err
	}

	err = WriteToClipboard(selection.Value)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	listCmd := ListCmd()
	rootCmd.AddCommand(listCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "search for a entry by keyword",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleSearchCommand(args)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleSearchCommand(args []string) (string, error) {

	if len(args) != 1 {
		return fmt.Sprintf("you need exactly one argument -> search query"), nil
	}

	closestN := 10

	result, err := Search(args[0], closestN)
	if err != nil {
		return "", err
	}

	output := ""
	if result == nil {
		output = "nix gefunden"
	} else {

		selection, err := ShowSelectionView(*result)
		if err != nil {
			return "", err
		}

		err = WriteToClipboard(selection.Value)
		if err != nil {
			return "", err
		}
	}

	return output, nil
}

func init() {
	searchCmd := SearchCmd()
	rootCmd.AddCommand(searchCmd)
}

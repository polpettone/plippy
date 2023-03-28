/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

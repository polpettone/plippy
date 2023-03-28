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

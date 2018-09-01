// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/FScoward/paper-sync/drobox"
	"github.com/spf13/cobra"
	"net/http"
)

// updateCmd represents the update command
func UpdateCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "update",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("update called")

			filePath, err := cmd.Flags().GetString("file_path")
			if err != nil {
				cmd.Println("Error", err)
				return
			}

			docId, err := cmd.Flags().GetString("doc_id")
			if err != nil || docId == "" {
				cmd.Println("Error", err)
				return
			}

			policy, err := cmd.Flags().GetString("doc_update_policy")
			if err != nil {
				cmd.Println("Error", err)
				return
			}

			revision, err := cmd.Flags().GetInt("revision")
			if err != nil {
				cmd.Println("Error", err)
				return
			}

			format, err := cmd.Flags().GetString("format")
			if err != nil {
				cmd.Println("Error", err)
				return
			}

			client := new(http.Client)
			_, err = drobox.UpdateDoc(client, filePath, docId, policy, revision, format)
			if err != nil {
				cmd.Println("Error", err)
			}

		},
	}
	command.Flags().StringP("file_path", "", "", "absolute file path")
	command.MarkFlagRequired("file_path")
	command.Flags().StringP("doc_id", "i", "", "document id")
	command.MarkFlagRequired("doc_id")
	command.Flags().StringP("doc_update_policy", "p", "", "document update policy")
	command.MarkFlagRequired("doc_update_policy")
	command.Flags().IntP("revision", "r", 0, "revision")
	command.MarkFlagRequired("revision")
	command.Flags().StringP("format", "f", "markdown", "format")
	return command
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

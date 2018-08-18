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
	"github.com/FScoward/paper-sync/drobox"
	"github.com/spf13/cobra"
	"net/http"
)

// downloadCmd represents the download command
func DownloadCmd() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			docId, err := cmd.Flags().GetString("doc_id")
			if err != nil || docId == "" {
				cmd.Println("Error", err)
				return
			}

			cmd.Println("download called", docId)
			client := new(http.Client)
			res, err := drobox.DownloadDoc(client, docId, "markdown")
			cmd.Println("// TITLE:", res.Header.Title)
			cmd.Println("// REVISION:", res.Header.Revision)
			cmd.Println("// MIME_TYPE:", res.Header.MimeType)
			cmd.Println("// OWNER:", res.Header.Owner)
			cmd.Println("")

			wantPreview, err := cmd.Flags().GetBool("preview")
			if err != nil {
				cmd.Println(err)
			}
			if wantPreview {
				cmd.Println(res.Body)
			}

			wantSave, err := cmd.Flags().GetBool("save")
			if err != nil {
				cmd.Println(err)
			}
			if wantSave {
				cmd.Println("Saving...")
				res.Save()
				cmd.Println("Done.")
			}
		},
	}
	downloadCmd.Flags().StringP("doc_id", "i", "", "document id")
	downloadCmd.Flags().BoolP("save", "s", false, "save?")
	downloadCmd.Flags().BoolP("preview", "p", false, "preview body?")
	return downloadCmd

}

/*func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "short",
		Long: `long`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print("version 0.1")
		},
	}
	return cmd
}*/

func init() {
	//rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

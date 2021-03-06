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
	"encoding/json"
	"fmt"
	"github.com/FScoward/paper-sync/drobox"
	"github.com/spf13/cobra"
	"net/http"
)

// paperListCmd represents the paperList command
func PaperListCmd() *cobra.Command {
	paperListCmd := &cobra.Command{

		Use:   "paperList",
		Short: "Show Paper List. ID and TITLE.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("paperList called\r")

			client := new(http.Client)
			idList := drobox.GetDocIdList(client)
			docIdList := &drobox.DocIdList{}
			json.Unmarshal(idList, &docIdList)
			responseList := map[string]drobox.DownloadDocResponse{}
			for _, v := range docIdList.DocIds {
				response, _ := drobox.DownloadDoc(client, v, "markdown")
				responseList[v] = response
				fmt.Printf("access... %v\r", v)
			}
			showResult(responseList)
		},
	}
	return paperListCmd
}

func showResult(responseList map[string]drobox.DownloadDocResponse) {
	for id, response := range responseList {
		fmt.Println("ID:", id, "TITLE:", response.Header.Title, "REVISION:", response.Header.Revision)
	}
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// paperListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// paperListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

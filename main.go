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

package main

import (
	"github.com/FScoward/paper-sync/cmd"
)

func main() {
	//client := new(http.Client)
	//fmt.Println(string(drobox.GetDocIdList(client)))
	cmd.Execute()
	//response, _ := drobox.DownloadDoc(client, "TTriQUccwfBTxlqEqSOlX", "markdown")
	//response.Save()
	//fmt.Println(response.Body)

	//_, err := drobox.UpdateDoc(client, "./README.md", "TTriQUccwfBTxlqEqSOlX", "overwrite_all", 48, "markdown")
	//fmt.Println(err)
}

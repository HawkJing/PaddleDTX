// Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package files

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	httpclient "github.com/PaddlePaddle/PaddleDTX/xdb/client/http"
)

var (
	replica int
)

// addNsCmd represents the command to add namespace
var addNsCmd = &cobra.Command{
	Use:   "addns",
	Short: "add file namespace into xuper db",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := httpclient.New(host)
		if err != nil {
			fmt.Printf("err：%v\n", err)
			return
		}
		if replica <= 0 || len(namespace) == 0 {
			fmt.Printf("err: bad param, replica and ns length must greater than 0")
			return
		}

		err = client.AddFileNs(context.Background(), privateKey, namespace, description, replica)
		if err != nil {
			fmt.Printf("err：%v\n", err)
			return
		}

		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(addNsCmd)

	addNsCmd.Flags().StringVarP(&privateKey, "privkey", "k", "", "private key")
	addNsCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace for file")
	addNsCmd.Flags().StringVarP(&description, "description", "d", "", "description")
	addNsCmd.Flags().IntVarP(&replica, "replica", "r", 0, "replica")

	addNsCmd.MarkFlagRequired("privkey")
	addNsCmd.MarkFlagRequired("namespace")
	addNsCmd.MarkFlagRequired("replica")
}

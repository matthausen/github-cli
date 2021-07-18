/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"net/http"
	"time"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Long: `Use your personal access token to delete a repository. Examples::

	github-cli delete --token <personal_access_token> <user_name/org_name> <repository_name>
	github-cli delete -t <personal_access_token> <user_name/org_name> <repository_name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")

		if token != "" {
			deleteRepo(args, token)
		} else {
			log.Fatal("You need a personal access token to delete a repository\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("token", "t", "", "Your personal access token")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// deleteRepo -> delete a user/org repository given a personal access token
func deleteRepo(args []string, token string){
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	req, err := http.NewRequest("DELETE", "https://api.github.com/repos/" + args[0] + "/" + args[1] , nil)
	req.Header.Set("Authorization", "token " + token)

	if err != nil {
		log.Fatalf("Could not delete the repository %s for user %s. Error: %v", args[0], args[1], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the delete api. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		fmt.Printf("Repository %s successfully deleted\n", args[1])
	} else {
		fmt.Printf("Could not delete repository. Response status code: %v\n", resp.StatusCode)
	}
}

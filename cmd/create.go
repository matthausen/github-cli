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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/matthausen/github-cli/model"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new repository",
	Long: `Create a new repository for a user or for an organization given a personal access token. For example:


github-cli create --token <personal_access_token> <repo_name>
github-cli create -t <personal_access_token> <repo_name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		isOrg, _ := cmd.Flags().GetBool("org")

		if token != "" {
			if isOrg {
				createOrgRepo(args, token)
			} else {
				createUserRepo(args, token)
			}
		} else {
			log.Fatal("You need a personal access token to create a new repository\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("org", "o", false, "Organization")
	createCmd.Flags().StringP("token", "t", "", "Personal access token")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// createOrgRepo ->  create a repository for an organization given a personal access token
func createOrgRepo(args []string, token string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	body := map[string]string{"name": args[0]}
	jsonValue, err := json.Marshal(body)

	if err!= nil {
		log.Fatalf("Could not marshal body of the request: %v\n", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/org/repos" , bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not create a new repository for this user. Error: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the create api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repo model.CreateRepoResponse

	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	fmt.Println(repo.Name, repo.Owner, repo.Private)
}

// createUserRepo -> create a repository for a user given a personal access token
func createUserRepo(args []string, token string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	body := map[string]string{"name": args[0]}
	jsonValue, err := json.Marshal(body)

	if err!= nil {
		log.Fatalf("Could not marshal body of the request: %v\n", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/user/repos" , bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not create a new repository for this user. Error: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the create api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repo model.CreateRepoResponse

	if err := json.NewDecoder(resp.Body).Decode(&repo); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	fmt.Println(repo.Name, repo.Owner, repo.Private)
}
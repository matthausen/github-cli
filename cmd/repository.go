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
	"encoding/json"
	"fmt"
	"github.com/matthausen/github-cli/model"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repos",
	Short: "Find the repository of a user or an organisation",
	Long: `Get all the repositories from a user profile or an organisation:

You can specify if you want to search an organisation by using the --org flag. E.g.:
github-cli repos <user_name>
github-cli repos --org <org_name>

You can also view private repos by passing the personal access token with the flag --token. E.g.:
github-cli repos --token <personal_access_token>
github-cli repos -o -t <personal_access_token>`,
	Run: func(cmd *cobra.Command, args []string) {
		isOrg, _ := cmd.Flags().GetBool("org")
		token, _ := cmd.Flags().GetString("token")

		if token != "" {
			if isOrg {
				fetchOrgReposWithCred(args, token)
			} else {
				fetchUserReposWithCred(args, token)
			}
		} else {
			if isOrg {
				fetchOrgRepos(args)
			} else {
				fetchUserRepos(args)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
	repositoryCmd.Flags().BoolP("org", "o", false, "Search for an organisation")
	repositoryCmd.Flags().StringP("token", "t", "", "Use personal access token to view private repos")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
// fetchOrgRepos -> fetch all repos given an organisation name
func fetchOrgRepos(args []string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/" + args[0] + "/repos?page=1&per_page=1000" , nil)

	if err != nil {
		log.Fatalf("Could not fetch repos from the organisation %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the orgs api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repos []model.Repository

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	for _, repo := range repos {
		fmt.Println(repo.HTMLUrl, repo.Language)
	}
	fmt.Printf("Total count: %d\n", len(repos))
}

// fetchUserRepos -> fetch all repos given a username
func fetchUserRepos(args []string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://api.github.com/users/" + args[0] + "/repos?page=1&per_page=1000" , nil)

	if err != nil {
		log.Fatalf("Could not fetch repos from the user %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the user api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repos []model.Repository

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	for _, repo := range repos {
		fmt.Println(repo.HTMLUrl, repo.Language)
	}

	fmt.Printf("Total count: %d\n", len(repos))
}

// fetchUserReposWithCred -> use personal access token to fetch public and private repos for a user
func fetchUserReposWithCred(args []string, token string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user/repos?page=1&per_page=1000" , nil)
	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not fetch repos from the user %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the user api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repos []model.Repository

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	for _, repo := range repos {
		fmt.Println(repo.HTMLUrl, repo.Language)
	}

	fmt.Printf("Total count: %d\n", len(repos))
}

// fetchOrgWithCred -> user personal access token to fetch public and private repos for an organization
func fetchOrgReposWithCred(args []string, token string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest("GET", "https://api.github.com/org/repos?page=1&per_page=1000" , nil)
	req.Header.Set("Authorization", "token " + token)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not fetch repos from the organisation %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the orgs api. Error: %v", err)
	}

	defer resp.Body.Close()

	var repos []model.Repository

	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		log.Fatalf("Could not decode response body for repositories. Error %v", err)
	}

	for _, repo := range repos {
		fmt.Println(repo.HTMLUrl, repo.Language)
	}
	fmt.Printf("Total count: %d\n", len(repos))
}
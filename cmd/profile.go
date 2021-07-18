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
	"log"
	"net/http"
	"time"

	"github.com/matthausen/github-cli/model"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Display public profile info",
	Long: `Get information about a profile or an org from GitHub:

You can specify if you want a user or an organization by using the flag --org. E.g.:
github-cli profile <name of the user>
github-cli profile --org <name of the org> or 
github-cli profile -o <name of the org>`,
	Run: func(cmd *cobra.Command, args []string) {
		org, _ := cmd.Flags().GetBool("org")

		if org {
			fetchOrg(args)
		} else {
			fetchUser(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.Flags().BoolP("org", "o", false, "Search for an org")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// fetchUser -> get public info of a user
func fetchUser(args []string){
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest("GET", "https://api.github.com/users/" + args[0], nil)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not fetch profile info for the user %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the user api. Error: %v", err)
	}

	defer resp.Body.Close()

	var user model.Profile

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatalf("Could not decode response body for user. Error %v", err)
	}

	fmt.Println(user.Name, user.Id, user.Avatar)
}

// fetchOrg -> get public information of an org
func fetchOrg(args []string) {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/" + args[0], nil)
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalf("Could not fetch profile info for the org %s. Error: %v", args[0], err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Could not get a response from the org api. Error: %v", err)
	}

	defer resp.Body.Close()

	var user model.Profile

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatalf("Could not decode response body for org. Error %v", err)
	}

	fmt.Println(user.Location, user.PublicRepos, user.TwitterHandle)
}

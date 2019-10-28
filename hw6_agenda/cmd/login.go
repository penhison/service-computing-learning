/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"agenda/entity"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login with given username and password",
	Long: `login with given username and password.
if username or password is incorrect, login will fail.
if you not register an account, use register instead.
`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		// fmt.Println("login called by " + username + " " + password)
		res := entity.Login(username, password)
		if !res.IsSuccess {
			fmt.Printf("login fail, reason: %s\n", res.Body)
		} else {
			fmt.Println(res.Body)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("user", "u", "Anonymous", "your username, if contain special characters, please quote it")
	loginCmd.Flags().StringP("password", "p", "", "your password, if contain special characters, please quote it")
	loginCmd.MarkFlagRequired("user")
	loginCmd.MarkFlagRequired("password")
}

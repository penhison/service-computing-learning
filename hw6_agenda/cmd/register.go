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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register with given username and password",
	Long: `register with given username and password.
username, password, email and telephone must be given,
if username is used by others, register will fail.
if you already register an account, use login instead.
`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")
		// fmt.Println("register called by " + username + " " + password + " " + email + " " + telephone)
		res := entity.Register(username, password, email, telephone)
		if !res.IsSuccess {
			fmt.Printf("register fail, reason: %s\n", res.Body)
		} else {
			fmt.Println(res.Body)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	registerCmd.Flags().StringP("user", "u", "", "your username, if contain special characters, please quote it")
	registerCmd.Flags().StringP("password", "p", "", "your password, if contain special characters, please quote it")
	registerCmd.Flags().StringP("email", "e", "", "your email, if contain special characters, please quote it")
	registerCmd.Flags().StringP("telephone", "t", "", "your telephone, if contain special characters, please quote it")
	registerCmd.MarkFlagRequired("user")
	registerCmd.MarkFlagRequired("password")
	registerCmd.MarkFlagRequired("email")
	registerCmd.MarkFlagRequired("telephone")
}
